package auth

import (
	"github.com/gofiber/fiber/v2"
)

type loginResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func LoginAttempt(c *fiber.Ctx) error {
	
	user := new(loginBody)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(loginResponse{fiber.StatusBadRequest, "Invalid Body Request"})
	}

	if (user.Username == "") || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(loginResponse{fiber.StatusBadRequest, "Missing username or password"})
	}

	if (user.Username == "admin" || user.Username == "admin@j4ck4l.com") && user.Password == "heyitsadmin" {
		return c.JSON(loginResponse{fiber.StatusOK, "Login successful"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(loginResponse{fiber.StatusUnauthorized, "Incorrect credentials"})
}


func RegisterAttempt(c *fiber.Ctx) error {
	
	user := new(loginBody)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(loginResponse{fiber.StatusBadRequest, "Invalid Body Request"})
	}

	if (user.Username == "") || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(loginResponse{fiber.StatusBadRequest, "Missing username or password"})
	}

	if (user.Username == "admin" || user.Username == "admin@j4ck4l.com") && user.Password == "heyitsadmin" {
		return c.JSON(loginResponse{fiber.StatusOK, "Login successful"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(loginResponse{fiber.StatusUnauthorized, "Incorrect credentials"})
}