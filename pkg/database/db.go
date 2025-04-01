package db

import (
	"fmt"
	"log"

	"github.com/j4ck4l-24/Ex0r/pkg/config"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

func InitDB() {
	postgresConfig, _, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	host = postgresConfig.Host
	port = postgresConfig.Port
	user = postgresConfig.Username
	password = postgresConfig.Password
	dbname = postgresConfig.DBname

	connString := fmt.Sprintf("%s %s %s %s %s", host, port, user, password, dbname)
	fmt.Println(connString)
}
