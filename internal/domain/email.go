package domain

import (
	"encoding/json"
	"log"
)

type BodyTemplateData struct {
	CompanyName string `json:"name"`
	Link        string `json:"link"`
}

type Email struct {
	To   []string
	Body string
}

type Payload struct {
	To           string           `json:"to"`
	TemplateName string           `json:"tn"`
	Data         BodyTemplateData `json:"d"`
}

func PayloadRedisToDTO(p string) *Payload {
	var pDto Payload
	err := json.Unmarshal([]byte(p), &pDto)
	if err != nil {
		log.Println("Error unmarshal Payload")
	}
	return &pDto
}
