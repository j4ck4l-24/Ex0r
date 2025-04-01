package config

import (
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Port     string
	Host     string
	Username string
	Password string
	DBname   string
}

type AppConfig struct {
	Port string
}

func Load() (*PostgresConfig, *AppConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, nil, err
	}

	postgresConfig := &PostgresConfig{
		Port:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB"),
	}

	appConfig := &AppConfig{
		Port: os.Getenv("APP_PORT"),
	}

	return postgresConfig, appConfig, nil
}
