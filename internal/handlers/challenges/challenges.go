package challenges

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
)

func GetAllChallenges(c *fiber.Ctx) error {
	dbConn, err := db.InitDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "DB Connection Failure. Contact Admin",
		})
	}
	role := c.Locals("user_data").(jwt.MapClaims)["role"].(string)

	if role != "admin" {
		var challenges []models.ChallengePublic

		query := "SELECT id, chall_name, chall_desc, category, current_points, max_attempts, type, author_name, connection_string, created_at, updated_at FROM Challenges WHERE hidden=FALSE"
		rows, err := dbConn.Query(query)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Failed to fetch challenges",
			})
		}
		defer rows.Close()
		for rows.Next() {
			var challenge models.ChallengePublic
			err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxAttempts, &challenge.Type, &challenge.AuthorName, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
					Status:  fiber.StatusInternalServerError,
					Message: "Something Went Wrong in fetching Challenges",
				})
			}
			challenges = append(challenges, challenge)
		}
		return c.Status(fiber.StatusOK).JSON(models.PublicChallengesResponse{
			Status:     fiber.StatusOK,
			Message:    "Challenges Fetched Succesfully",
			Challenges: challenges,
		})
	} else {
		query := "SELECT id, chall_name, chall_desc, category, current_points, initial_points, min_points, max_attempts, type, hidden, author_name, decay_type, decay_value, connection_string, created_at, updated_at /*,requirements, next_chall_id*/ FROM Challenges "
		rows, err := dbConn.Query(query)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Failed to fetch challenges",
			})
		}
		defer rows.Close()
		var challenges []models.ChallengeAdmin
		for rows.Next() {
			var challenge models.ChallengeAdmin
			err = rows.Scan(&challenge.ChallId, &challenge.ChallName, &challenge.ChallDesc, &challenge.Category, &challenge.Points, &challenge.MaxPoints, &challenge.MinPoints, &challenge.MaxAttempts, &challenge.Type, &challenge.Hidden, &challenge.AuthorName, &challenge.DecayType, &challenge.DecayValue, &challenge.ConnectionString, &challenge.CreatedAt, &challenge.UpdatedAt)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
					Status:  fiber.StatusInternalServerError,
					Message: "Something Went Wrong in fetching Challenges",
				})
			}
			challenges = append(challenges, challenge)
		}
		return c.Status(fiber.StatusOK).JSON(models.AdminChallengesResponse{
			Status:     fiber.StatusOK,
			Message:    "Challenges Fetched Succesfully",
			Challenges: challenges,
		})
	}
}
