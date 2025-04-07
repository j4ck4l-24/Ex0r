package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/j4ck4l-24/Ex0r/pkg/database"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type successfulLoginResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Token string `json:"token"`
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var dbConn, err = db.InitDB()

func LoginAttempt(c *fiber.Ctx) error {

	user := new(loginBody)
	var userId int
	var storedHashPassword string
	var userEmail string
	var userRole string
	var dbUserName string
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{fiber.StatusBadRequest, "Invalid Body Request"})
	}

	if (user.Username == "") || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{fiber.StatusBadRequest, "Missing username or password"})
	}

	query := "SELECT id, username, email, role, password_hash FROM Users WHERE username = $1"
	err = dbConn.QueryRow(query, user.Username).Scan(&userId, &userEmail, &dbUserName, &userRole, &storedHashPassword)

	if err != nil {
		// fmt.Print("User not found: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(response{fiber.StatusUnauthorized, "Incorrect credentials"})
	}
	if !utils.VerifyPassword(user.Password, storedHashPassword) {
		return c.Status(fiber.StatusUnauthorized).JSON(response{fiber.StatusUnauthorized, "Incorrect credentials"})
	}
	token, err := utils.CreateToken(userId, dbUserName, userEmail, userRole)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(response{fiber.StatusBadGateway, "Something Went Wrong"})
	}
	// fmt.Printf("%v:%v", userId, storedHashPassword)
	return c.Status(fiber.StatusOK).JSON(successfulLoginResponse{fiber.StatusOK, "Login Successful", token})
}

func RegisterAttempt(c *fiber.Ctx) error {

	user := new(registerBody)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{fiber.StatusBadRequest, "Invalid Body Request"})
	}

	if (user.Username == "") || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{fiber.StatusBadRequest, "Missing username or password"})
	}

	if !utils.IsValidEmail(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(response{fiber.StatusBadRequest, "Invalid email"})
	}
	user.Password, _ = utils.HashPassword(user.Password)
	var userId int
	query := "INSERT INTO USERS (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id"

	err = dbConn.QueryRow(query, user.Username, user.Password, strings.ToLower(user.Email)).Scan(&userId)
	if err != nil {
		// fmt.Print("Already exists a User with same name or email", err)
		return c.Status(fiber.StatusUnauthorized).JSON(response{fiber.StatusUnauthorized, "User already exists"})
	}

	return c.Status(fiber.StatusOK).JSON(response{fiber.StatusOK, "Registration Successful"})
}
