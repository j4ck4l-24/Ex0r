package config

import (
	"os"

	"github.com/j4ck4l-24/Ex0r/pkg/models"
	"github.com/joho/godotenv"
)

func Load() (*models.PostgresConfig, *models.AppConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, nil, err
	}

	postgresConfig := &models.PostgresConfig{
		Port:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB"),
	}

	appConfig := &models.AppConfig{
		Port:      os.Getenv("APP_PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	return postgresConfig, appConfig, nil
}
