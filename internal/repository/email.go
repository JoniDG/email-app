package repository

import (
	"email-app/internal/defines"
	"email-app/internal/domain"
	"log"
	"net/smtp"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type EmailRepository interface {
	SendMail(email domain.Email) error
}

type emailRepository struct {
	auth smtp.Auth
	host string
	from string
}

func NewEmailRepository(auth smtp.Auth) EmailRepository {
	return &emailRepository{
		auth: auth,
	}
}

func (r *emailRepository) SendMail(email domain.Email) error {
	addr := os.Getenv(defines.EnvEmailHost) + ":" + os.Getenv(defines.EnvEmailPort)
	body := []byte(email.Body)
	log.Printf("Body length %d bytes\n", len(body))
	err := smtp.SendMail(addr, r.auth, os.Getenv("EMAIL_SENDER_USER"), email.To, body)
	if err != nil {
		return err
	} else {
		log.Println("Email Enviado")
		return nil
	}
}
