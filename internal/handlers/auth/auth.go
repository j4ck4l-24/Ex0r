package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func LoginAttempt(c *fiber.Ctx) error {

	user := new(models.LoginBody)
	var userId int
	var storedHashPassword string
	var userEmail string
	var userRole string
	var dbUserName string
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid Body Request",
		})
	}

	if (user.Username == "") || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Missing username or password",
		})
	}
	var dbConn, err = db.InitDB()

	query := "SELECT id, username, email, role, password_hash FROM Users WHERE username = $1"
	err = dbConn.QueryRow(query, user.Username).Scan(&userId, &dbUserName, &userEmail, &userRole, &storedHashPassword)

	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(models.GeneralResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Incorrect credentials",
		})
	}
	if !utils.VerifyPassword(user.Password, storedHashPassword) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Incorrect credentials",
		})
	}
	token, err := utils.CreateToken(userId, dbUserName, userEmail, userRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Something Went Wrong",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.SuccessfulLoginResponse{
		Status:  fiber.StatusOK,
		Message: "Login Successful",
		Token:   token,
	})
}

func RegisterAttempt(c *fiber.Ctx) error {

	user := new(models.RegisterBody)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid Body Request",
		})
	}

	if (user.Username == "") || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Missing username or password",
		})
	}

	if !utils.IsValidEmail(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid email",
		})
	}
	var dbConn, err = db.InitDB()
	user.Password, _ = utils.HashPassword(user.Password)
	var userId int
	query := "INSERT INTO USERS (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id"

	err = dbConn.QueryRow(query, user.Username, user.Password, strings.ToLower(user.Email)).Scan(&userId)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "User already exists",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "Registration Successful",
	})
}
