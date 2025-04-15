package challenges

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/services"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func GetAllChallenges(c *fiber.Ctx) error {
	dbConn := db.DB
	role := c.Locals("user_data").(jwt.MapClaims)["role"].(string)
	if role != "Admin" {
		challenges, err := services.FetchPublicChallenges(c, dbConn)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to fetch challenges")
		}
		if len(challenges) == 0 {
			return utils.SendGeneralResp(c, fiber.StatusNotFound, "No challenges found")
		}
		return utils.SendMultiplePublicChallsResp(c, fiber.StatusOK, "Challenges Fetched Succesfully", challenges)

	} else {
		challenges, err := services.FetchAdminChallenges(c, dbConn)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to fetch challenges")
		}
		if len(challenges) == 0 {
			return utils.SendGeneralResp(c, fiber.StatusNotFound, "No challenges found")
		}
		return utils.SendMultipleAdminChallsResp(c, fiber.StatusOK, "Challenges Fetched Succesfully", challenges)
	}
}

func GetChallenge(c *fiber.Ctx) error {
	dbConn := db.DB
	role := c.Locals("user_data").(jwt.MapClaims)["role"].(string)

	if role != "Admin" {
		challenge, err := services.FetchPublicChallenge(c, dbConn)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid chall_id")
		}
		return utils.SendSinglePublicChallResp(c, fiber.StatusOK, "Challenge fetched Succesfully", challenge)
	} else {
		challenge, err := services.FetchAdminChallenge(c, dbConn)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid chall_id")
		}
		return utils.SendSingleAdminChallResp(c, fiber.StatusOK, "Challenge fetched Succesfully", challenge)
	}
}

func CreateChallenge(c *fiber.Ctx) error {
	dbConn := db.DB
	challenge := new(models.AdminChallenge)

	if err := c.BodyParser(challenge); err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if challenge.Type == "static" {
		query := "INSERT INTO Challenges (chall_name, chall_desc, category, current_points, initial_points, min_points, max_attempts, type, hidden, author_name, connection_string) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
		err := dbConn.QueryRow(query, challenge.ChallName, challenge.ChallDesc, challenge.Category, challenge.Points, challenge.MaxPoints, challenge.MinPoints, challenge.MaxAttempts, challenge.Type, challenge.Hidden, challenge.AuthorName, challenge.ConnectionString).Scan(&challenge.ChallId)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to create challenge")
		}
		return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Challenge created successfully with id = %d", challenge.ChallId))

	} else if challenge.Type == "dynamic" {
		query := "INSERT INTO Challenges (chall_name, chall_desc, category, current_points, initial_points, min_points, max_attempts, type, hidden, author_name, decay_type, decay_value, connection_string) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id"
		err := dbConn.QueryRow(query, challenge.ChallName, challenge.ChallDesc, challenge.Category, challenge.Points, challenge.MaxPoints, challenge.MinPoints, challenge.MaxAttempts, challenge.Type, challenge.Hidden, challenge.AuthorName, challenge.DecayType, challenge.DecayValue, challenge.ConnectionString).Scan(&challenge.ChallId)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to create challenge")
		}
		return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Challenge created successfully with id = %d", challenge.ChallId))
	}

	return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid challenge type")
}

func UpdateChallenge(c *fiber.Ctx) error {
	dbConn := db.DB
	challenge := new(models.AdminChallenge)
	challenge.ChallId, _ = c.ParamsInt("id")
	if err := c.BodyParser(challenge); err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if challenge.Type == "static" {
		challenge, err := services.UpdateStaticChallenge(dbConn, challenge)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to update challenge either there is an error or challenge with chall_id doens't exist")
		}
		return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Challenge updated successfully with id = %d", challenge.ChallId))

	} else if challenge.Type == "dynamic" {
		challenge, err := services.UpdateDynamicChallenge(dbConn, challenge)
		if err != nil {
			return utils.SendGeneralResp(c, fiber.StatusInternalServerError, "Failed to update challenge either there is an error or challenge with chall_id doens't exist")
		}
		return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Challenge updated successfully with id = %d", challenge.ChallId))
	}
	return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid challenge type")
}

func DeleteChallenge(c *fiber.Ctx) error {
	dbConn := db.DB
	chall_id, _ := c.ParamsInt("id")

	query := "DELETE FROM Challenges WHERE id = $1 RETURNING id"

	err := dbConn.QueryRow(query, chall_id).Scan(&chall_id)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusOK, "Failed to delete challenge")
	}
	return utils.SendGeneralResp(c, fiber.StatusOK, fmt.Sprintf("Challenge deleted successfully with chall_id = %d", chall_id))
}
