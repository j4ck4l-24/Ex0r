package utils

import (
	"fmt"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(id int, username string, email string, role string) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  id,
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err = t.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(token string) (claims jwt.MapClaims, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Method)
		}
		return SECRET_KEY, nil
	})

	fmt.Printf("%v \n %v \n",t.Method ,t.Valid)
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
