package main

import (
	"net/http"
    "log"

    "go-data/db"
    "go-data/handlers"
)

func main() {
    db.Init()

    http.HandleFunc("/users", handlers.CreateUser)
    http.HandleFunc("/verify/", handlers.ConfirmEmail)
    log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
