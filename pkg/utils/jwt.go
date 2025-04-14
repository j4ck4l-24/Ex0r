package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
	if token == "" {
		return nil, fmt.Errorf("empty token")
	}
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Method)
		}
		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func IsAdmin(token string) (isAdmin bool) {
	t, _ := VerifyToken(token)
	if t == nil {
		return false
	}
	return t["role"] == "Admin"
}

func ValidToken(token string) (validToken bool) {
	t, _ := VerifyToken(token)

	return t != nil
}
