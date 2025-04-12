package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/j4ck4l-24/Ex0r/pkg/utils"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" || !utils.ValidToken(token) {
			return c.Status(fiber.StatusUnauthorized).JSON(models.GeneralResponse{
				Status: fiber.StatusUnauthorized, Message: "Unauthorized Access",
			})
		}
		claims, _ := utils.VerifyToken(token)
		c.Locals("user_data", claims)
		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if !utils.IsAdmin(token) {
			return c.Status(fiber.StatusForbidden).JSON(models.GeneralResponse{
				Status:  fiber.StatusForbidden,
				Message: "Forbidden: Admin access required",
			})
		}
		return c.Next()
	}
}
