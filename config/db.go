package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	if err := godotenv.Load(".env.test"); err != nil {
		_ = godotenv.Load()
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return DB.Ping()
}

func CreateApplicationTable() {
	query := `
	CREATE TABLE IF NOT EXISTS applications (
		id SERIAL PRIMARY KEY,
		full_name TEXT,
		email TEXT,
		dob DATE,
		program_applied TEXT
	)`
	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("❌ Failed to create applications table: %v", err)
	}
}
