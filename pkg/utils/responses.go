package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
)

func SendGeneralResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(models.GeneralResponse{
		Status:  statusCode,
		Message: message,
	})
}

func SendPublicChallengeResponse(c *fiber.Ctx, statusCode int, message string, challenge models.ChallengePublic) error {
	response := models.PublicChallengeResponse{
		Status:    statusCode,
		Message:   message,
		Challenge: challenge,
	}
	return c.Status(statusCode).JSON(response)
}
func SendAdminChallengeResponse(c *fiber.Ctx, statusCode int, message string, challenge models.ChallengeAdmin) error {
	response := models.AdminChallengeResponse{
		Status:    statusCode,
		Message:   message,
		Challenge: challenge,
	}
	return c.Status(statusCode).JSON(response)
}

func SendPublicChallengesResponse(c *fiber.Ctx, statusCode int, message string, challenges []models.ChallengePublic) error {
	response := models.PublicChallengesResponse{
		Status:     statusCode,
		Message:    message,
		Challenges: challenges,
	}
	return c.Status(statusCode).JSON(response)
}

func SendAdminChallengesResponse(c *fiber.Ctx, statusCode int, message string, challenges []models.ChallengeAdmin) error {
	response := models.AdminChallengesResponse{
		Status:     statusCode,
		Message:    message,
		Challenges: challenges,
	}
	return c.Status(statusCode).JSON(response)
}

func SendSuccessfulLoginResponse(c *fiber.Ctx, token string) error {
	response := models.SuccessfulLoginResponse{
		Status:  fiber.StatusOK,
		Message: "Login Successful",
		Token:   token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
