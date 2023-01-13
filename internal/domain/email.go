package domain

type BodyTemplateData struct {
	Company string
	Link    string
}

type Email struct {
	To      []string
	Subject string
	Body    string
}

type Payload struct {
	To           []string
	NameTemplate string
	Data         BodyTemplateData
}
