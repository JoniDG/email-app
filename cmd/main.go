package main

import (
	"email-app/internal/defines"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/smtp"
	"os"
)

func main() {
	fmt.Println("Email running")
	email()
}

func email() {
	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")

	toEmail := defines.To
	to := []string{toEmail}

	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port
	subject := defines.TitleEmail
	body := defines.BodyEmail
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Enviado")
}
