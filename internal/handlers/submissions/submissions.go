package submissions

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/services"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func HitSubmission(c *fiber.Ctx) error {
	ip := c.IP()
	submission := new(models.Submission)
	userData, ok := c.Locals("user_data").(jwt.MapClaims)
	if !ok {
		return utils.SendGeneralResp(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	dbConn := db.DB
	if err := c.BodyParser(submission); err != nil || submission.ChallId == 0 || submission.Submitted == "" {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	submission.UserId = int(userData["user_id"].(float64))
	submission.TeamId = int(userData["team_id"].(float64))

	if services.CheckSolved(submission) {
		return utils.SendGeneralResp(c, fiber.StatusConflict, "Already Solved")
	}
	query := "INSERT INTO Submissions (submitted, chall_id, user_id, team_id, ip) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at"
	err := dbConn.QueryRow(query, submission.Submitted, submission.ChallId, submission.UserId, submission.TeamId, ip).Scan(&submission.SubmissionId, &submission.CreatedAt, &submission.UpdatedAt)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Something Went Wrong")
	}

	if services.CheckForCorrectSubmission(submission) {
		err := services.UpdateSolves(submission)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Something Went Wrong")
		}
		return utils.SendSubmissionStatusResp(c, fiber.StatusOK, "Correct Submission", true)
	}
	return utils.SendSubmissionStatusResp(c, fiber.StatusOK, "Wrong Submission", false)

}

func GetAllSubmissions(c *fiber.Ctx) error {
	m := c.Queries()
	args := []interface{}{}
	idx := 1
	dbConn := db.DB
	query := "SELECT id, submitted, chall_id, user_id, team_id, ip, created_at, updated_at FROM Submissions"
	conditions := []string{}

	if strChall_id := m["chall_id"]; strChall_id != "" && strChall_id != "null" {
		if chall_id, err := strconv.Atoi(strChall_id); err == nil && chall_id > 0 {
			conditions = append(conditions, fmt.Sprintf("chall_id = $%d", idx))
			args = append(args, chall_id)
			idx++
		}
	}

	if strUser_id := m["user_id"]; strUser_id != "" && strUser_id != "null" {
		if user_id, err := strconv.Atoi(strUser_id); err == nil && user_id > 0 {
			conditions = append(conditions, fmt.Sprintf("user_id = $%d", idx))
			args = append(args, user_id)
			idx++
		}
	}

	if strTeam_id := m["team_id"]; strTeam_id != "" && strTeam_id != "null" {
		if team_id, err := strconv.Atoi(strTeam_id); err == nil && team_id > 0 {
			conditions = append(conditions, fmt.Sprintf("team_id = $%d", idx))
			args = append(args, team_id)
			idx++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	limit, offset := utils.Pagination(c)

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, offset)

	rows, err := dbConn.Query(query, args...)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Submissions not found")
	}
	defer rows.Close()
	var submissions []models.Submission
	for rows.Next() {
		var submission models.Submission
		err = rows.Scan(&submission.SubmissionId, &submission.Submitted, &submission.ChallId, &submission.UserId, &submission.TeamId, &submission.Ip, &submission.CreatedAt, &submission.UpdatedAt)
		if err != nil {
			log.Print(err)
			return utils.SendGeneralResp(c, fiber.StatusNotFound, "Submissions not found")
		}
		submissions = append(submissions, submission)
	}

	if len(submissions) == 0 {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Submissions not found")
	}
	return utils.SendMultipleSubmissionResp(c, fiber.StatusOK, "Submissions fetched Successfully", submissions)

}

func GetSubmission(c *fiber.Ctx) error {
	dbConn := db.DB
	submission_id, err := c.ParamsInt("id")
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Invalid submission_id")
	}
	query := "SELECT id, submitted, chall_id, user_id, team_id, ip, created_at, updated_at FROM Submissions WHERE id = $1"

	rows, err := dbConn.Query(query, submission_id)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Submission not found")
	}
	defer rows.Close()
	var submission models.Submission
	rows.Next()
	err = rows.Scan(&submission.SubmissionId, &submission.Submitted, &submission.ChallId, &submission.UserId, &submission.TeamId, &submission.Ip, &submission.CreatedAt, &submission.UpdatedAt)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Submission not found")
	}
	return utils.SendSingleSubmissionResp(c, fiber.StatusOK, "Submission fetched Successfully", submission)
}

func UpdateSubmission(c *fiber.Ctx) error {
	dbConn := db.DB
	submission := new(models.Submission)
	err := c.BodyParser(submission)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	submission.SubmissionId, err = c.ParamsInt("id")
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusNotFound, "Invalid submission_id")
	}

	query := "UPDATE Submissions SET submitted = $1, chall_id = $2, user_id = $3, team_id = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id;"
	err = dbConn.QueryRow(query, submission.Submitted, submission.ChallId, submission.UserId, submission.TeamId, submission.SubmissionId).Scan(&submission.SubmissionId)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to update submission either there is an error or submission with submission_id doens't exist")
	}
	return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Submission updated successfully with id = %d", submission.SubmissionId))
}

func DeleteSubmission(c *fiber.Ctx) error {
	dbConn := db.DB
	submission_id, _ := c.ParamsInt("id")

	query := "DELETE FROM Submissions WHERE id = $1 RETURNING id"

	err := dbConn.QueryRow(query, submission_id).Scan(&submission_id)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusOK, "Failed to delete submission")
	}
	return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Submission deleted successfully with chall_id = %d", submission_id))
}
