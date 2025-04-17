package submissions

import (
	"log"

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
	userData := c.Locals("user_data").(jwt.MapClaims)
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
			log.Print(err.Error())
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Something Went Wrong")
		}
	}
	return utils.SendSingleSubmissionResp(c, fiber.StatusOK, "Submissited Succesfully", *submission)
}
