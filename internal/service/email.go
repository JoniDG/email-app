package service

import (
	"bytes"
	"email-app/internal/defines"
	"email-app/internal/domain"
	"email-app/internal/repository"
	"html/template"
	"path/filepath"
)

type EmailService interface {
	ParseTemplate(payload domain.Payload) error
}

type emailService struct {
	emailRepo repository.EmailRepository
}

func NewEmailService(emailRepo repository.EmailRepository) EmailService {
	return &emailService{
		emailRepo: emailRepo,
	}
}

func (s *emailService) ParseTemplate(payload domain.Payload) error {
	fp := filepath.Join("./templates", payload.TemplateName+".html")
	t, err := template.ParseFiles(fp)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, payload.Data)
	if err != nil {
		return err
	}
	email := domain.Email{
		To: []string{payload.To},
		Body: domain.BodyMail{
			Headers: "From: " + payload.Data.CompanyName + "\n" +
				"To: " + payload.To + "\n" +
				defines.SubjectEmail + "\n" +
				defines.Mime + "\r\n",
			Message: buf.String(),
		},
	}
	err = s.emailRepo.SendMail(email)
	if err != nil {
		return err
	}
	return nil
}
