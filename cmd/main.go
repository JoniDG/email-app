package main

import (
	"bytes"
	"email-app/internal/defines"
	"email-app/internal/domain"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"net/smtp"
	"os"
	"time"
)

func main() {
	fmt.Println("Email running")
	email()
}

func email() {
	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	auth := smtp.PlainAuth("", from, password, host)
	address := host + ":" + port
	toEmail := defines.To
	to := []string{toEmail}
	t, err := template.ParseFiles("./templates/body.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	msg := domain.Body{
		Name:        "Jonathan",
		TimeStamp:   time.Now(),
		Temperature: 30,
	}
	err = t.Execute(buf, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := defines.TitleEmail + "\n"
	body += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";" + "\r\n"
	body += buf.String()
	err = smtp.SendMail(address, auth, from, to, []byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Enviado")
}
