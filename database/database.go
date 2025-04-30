package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var connStr = "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable"
var DB *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("✅ Connected to database")
	return db
}

func RunSQLFile(db *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read SQL file: %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL file: %w", err)
	}

	return nil

}
