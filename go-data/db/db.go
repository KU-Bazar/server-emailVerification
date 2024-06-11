package db

import (
    "log"

    "database/sql"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
    var err error
    dsn := "user=postgres password=#Redcarpet2552 dbname=new sslmode=disable"
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

