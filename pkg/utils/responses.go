package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j4ck4l-24/Ex0r/pkg/models"
)

func SendGeneralResp(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(models.GeneralResponse{
		Status:  statusCode,
		Message: message,
	})
}

func SendSinglePublicChallResp(c *fiber.Ctx, statusCode int, message string, challenge models.PublicChallenge) error {
	response := models.SinglePublicChallResp{
		Status:  statusCode,
		Message: message,
		Data:    challenge,
	}
	return c.Status(statusCode).JSON(response)
}
func SendSingleAdminChallResp(c *fiber.Ctx, statusCode int, message string, challenge models.AdminChallenge) error {
	response := models.SingleAdminChallResp{
		Status:  statusCode,
		Message: message,
		Data:    challenge,
	}
	return c.Status(statusCode).JSON(response)
}

func SendMultiplePublicChallsResp(c *fiber.Ctx, statusCode int, message string, challenges []models.PublicChallenge) error {
	response := models.PublicChallengesResp{
		Status:  statusCode,
		Message: message,
		Data:    challenges,
	}
	return c.Status(statusCode).JSON(response)
}

func SendMultipleAdminChallsResp(c *fiber.Ctx, statusCode int, message string, challenges []models.AdminChallenge) error {
	response := models.AdminChallengesResp{
		Status:  statusCode,
		Message: message,
		Data:    challenges,
	}
	return c.Status(statusCode).JSON(response)
}

func SendLoginSuccessResp(c *fiber.Ctx, token string) error {
	response := models.SuccessfulLoginResponse{
		Status:  fiber.StatusOK,
		Message: "Login Successful",
		Token:   token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func SendSingleFlagResp(c *fiber.Ctx, statusCode int, message string, flag models.Flag) error {
	response := models.SingleFlagResp{
		Status:  statusCode,
		Message: message,
		Data:    flag,
	}
	return c.Status(statusCode).JSON(response)
}

func SendMultipleFlagResp(c *fiber.Ctx, statusCode int, message string, flags []models.Flag) error {
	response := models.MultipleFlagResp{
		Status:  statusCode,
		Message: message,
		Data:    flags,
	}
	return c.Status(statusCode).JSON(response)
}

func SendSingleSubmissionResp(c *fiber.Ctx, statusCode int, message string, submission models.Submission) error {
	response := models.SingleSubmissionResp{
		Status:  statusCode,
		Message: message,
		Data:    submission,
	}
	return c.Status(statusCode).JSON(response)
}

func SendMultipleSubmissionResp(c *fiber.Ctx, statusCode int, message string, submissions []models.Submission) error {
	response := models.MultipleSubmissionResp{
		Status:  statusCode,
		Message: message,
		Data:    submissions,
	}
	return c.Status(statusCode).JSON(response)
}

func SendSubmissionStatusResp(c *fiber.Ctx, statusCode int, message string, isCorrect bool) error {
	response := models.SubmissionStatusResp{
		Status:    statusCode,
		Message:   message,
		IsCorrect: isCorrect,
	}
	return c.Status(statusCode).JSON(response)
}
