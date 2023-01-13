package repository

import (
	"email-app/internal/domain"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/smtp"
	"os"
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
	addr := os.Getenv("EMAIL_HOST") + ":" + os.Getenv("EMAIL_PORT")
	err := smtp.SendMail(addr, r.auth, os.Getenv("EMAIL_SENDER_USER"), email.To, []byte(email.Body))
	if err != nil {
		return err
	} else {
		log.Println("Email Enviado")
		return nil
	}
}
