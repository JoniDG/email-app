package main

import (
	"email-app/internal/controller"
	"email-app/internal/defines"
	"email-app/internal/domain"
	"email-app/internal/repository"
	"email-app/internal/service"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/smtp"
	"os"
)

func main() {
	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	auth := smtp.PlainAuth("", from, password, host)
	emailRepo := repository.NewEmailRepository(auth)
	svc := service.NewEmailService(emailRepo)
	ctrl := controller.NewEmailController(svc)

	msg := domain.BodyTemplateData{
		Company: "Example Company",
		Link:    "http://example.com",
	}

	payload := domain.Payload{
		To:           []string{defines.To},
		NameTemplate: defines.NameTemplate,
		Data:         msg,
	}
	fmt.Println("Email running")
	ctrl.Handle(payload)
}
