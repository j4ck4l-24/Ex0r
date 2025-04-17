package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func LoginAttempt(c *fiber.Ctx) error {

	reqBody := new(models.LoginBody)
	var user models.User
	if err := c.BodyParser(reqBody); err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if (reqBody.Username == "") || reqBody.Password == "" {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Missing username or password")
	}
	dbConn := db.DB

	query := "SELECT id, username, email, role, COALESCE(team_id, 0), password_hash FROM Users WHERE username = $1"
	err := dbConn.QueryRow(query, reqBody.Username).Scan(&user.UserId, &user.DbUserName, &user.UserEmail, &user.UserRole, &user.TeamId, &user.StoredHashPassword)

	if err != nil || !utils.VerifyPassword(reqBody.Password, user.StoredHashPassword) {
		return utils.SendGeneralResp(c, fiber.StatusUnauthorized, "Incorrect credentials")
	}

	token, err := utils.CreateToken(user.UserId, user.DbUserName, user.UserEmail, user.UserRole, user.TeamId)

	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Something Went Wrong")
	}
	return utils.SendLoginSuccessResp(c, token)
}

func RegisterAttempt(c *fiber.Ctx) error {

	reqBody := new(models.RegisterBody)

	if err := c.BodyParser(reqBody); err != nil {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if (reqBody.Username == "") || reqBody.Email == "" || reqBody.Password == "" {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Missing username or password")
	}

	if !utils.IsValidEmail(reqBody.Email) {
		return utils.SendGeneralResp(c, fiber.StatusBadRequest, "Invalid Email")
	}
	dbConn := db.DB
	reqBody.Password, _ = utils.HashPassword(reqBody.Password)
	var userId int
	query := "INSERT INTO USERS (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id"

	err := dbConn.QueryRow(query, reqBody.Username, reqBody.Password, strings.ToLower(reqBody.Email)).Scan(&userId)
	if err != nil {
		return utils.SendGeneralResp(c, fiber.StatusConflict, "User Already Exists")
	}

	return utils.SendGeneralResp(c, fiber.StatusOK, "Registration Successful")

}
