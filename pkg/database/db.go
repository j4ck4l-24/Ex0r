package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/j4ck4l-24/Ex0r/pkg/config"
	_ "github.com/lib/pq"
)
var DB *sql.DB

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

func InitDB() error {
	postgresConfig, _, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load database config: %v", err)
	}

	host = postgresConfig.Host
	port = postgresConfig.Port
	user = postgresConfig.Username
	password = postgresConfig.Password
	dbname = postgresConfig.DBname

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Printf("Error connecting the db:%v", err)
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)                 
	db.SetConnMaxLifetime(10 * time.Minute) 

	err = db.Ping()
	if err != nil {
		fmt.Print("Error pinging the db")
		return err
	}
	DB = db
	return nil
}
