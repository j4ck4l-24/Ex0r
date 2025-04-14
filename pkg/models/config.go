package models

type PostgresConfig struct {
	Port     string
	Host     string
	Username string
	Password string
	DBname   string
}

type AppConfig struct {
	Port      string
	JwtSecret string
}
