package config

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// LoadEnv loads variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}
}
