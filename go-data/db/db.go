package db

import (
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {

    err_env := godotenv.Load()
    if err_env != nil {
        log.Fatal("Error loading .env file")
    }
    
    var err error
    dsn := os.Getenv("DSN")
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

