package controller

import (
	"email-app/internal/domain"
	"email-app/internal/service"
	"log"
)

type EmailController interface {
	Handle(payload *string)
}

type emailController struct {
	svc service.EmailService
}

func NewEmailController(svc service.EmailService) EmailController {
	return &emailController{
		svc: svc,
	}
}

func (e *emailController) Handle(payload *string) {
	pDTO := domain.PayloadRedisToDTO(*payload)
	err := e.svc.ParseTemplate(*pDTO)
	if err != nil {
		log.Printf("Error Parsing Template/Data")
	}
}
