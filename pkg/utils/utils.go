package utils

import (
	"net/mail"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Pagination(c *fiber.Ctx) (limit int, offset int) {
	m := c.Queries()
	limit = 10
	page := 1
	if strLimit := m["limit"]; strLimit != "" && strLimit != "null" {
		if parsedLimit, err := strconv.Atoi(strLimit); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if pageStr := m["page"]; pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	offset = (page -1) * limit
	return limit, offset
}
