package controller

import (
	"email-app/internal/domain"
	"email-app/internal/service"
	"log"
)

type EmailController interface {
	Handle(payload domain.Payload)
}

type emailController struct {
	svc service.EmailService
}

func NewEmailController(svc service.EmailService) EmailController {
	return &emailController{
		svc: svc,
	}
}

func (e *emailController) Handle(payload domain.Payload) {
	err := e.svc.ParseTemplate(payload)
	if err != nil {
		log.Printf("Error Parsing Template/Data")
	}
}
