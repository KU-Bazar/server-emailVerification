package data

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)


func GenerateToken() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func SendVerificationEmail(email, token string) error {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	from := os.Getenv("email")
	password := os.Getenv("password")
	to := email
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "Subject: Email Verification\n\n" +
		"Please verify your email by clicking the following link:\n" +
		fmt.Sprintf("http://localhost:8080/verify/%s", token)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err_smpt := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	return err_smpt
}



