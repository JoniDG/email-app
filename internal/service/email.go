package service

import (
	"bytes"
	"email-app/internal/defines"
	"email-app/internal/domain"
	"email-app/internal/repository"
	"html/template"
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
	t, err := template.ParseFiles(payload.NameTemplate)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, payload.Data)
	if err != nil {
		return err
	}
	email := domain.Email{
		To:      payload.To,
		Subject: defines.SubjectEmail,
		Body:    defines.SubjectEmail + "\n" + defines.Mime + buf.String(),
	}
	err = s.emailRepo.SendMail(email)
	if err != nil {
		return err
	}
	return nil
}
