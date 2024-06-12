package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"go-data/db"
	"go-data/data"
	"go-data/models"
)

var (
	verificationStore = make(map[string]string)
	mu sync.Mutex
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    
	var req models.User

    fmt.Println("2")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
    fmt.Println("3")
	token := data.GenerateToken()
	mu.Lock()
	verificationStore[token] = req.Email
	mu.Unlock()
    fmt.Println("4")
	err = data.SendVerificationEmail(req.Email, token)
	if err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}
    fmt.Println("5")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Verification email sent"))
 
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Path[len("/verify/"):]
	mu.Lock()
	email, exists := verificationStore[token]
	if exists {
		delete(verificationStore, token)
	}
	mu.Unlock()

	if !exists {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	fmt.Println("Email verified:", email)

	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		username VARCHAR(255) NOT NULL
	);
	`
	_, err_query := db.DB.Exec(query)
		if err_query != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email verified successfully"))
}

