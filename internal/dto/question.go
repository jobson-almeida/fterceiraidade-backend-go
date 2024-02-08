package dto

import (
	"github.com/lib/pq"
)

type QuestionInput struct {
	ID           string         `json:"id"`
	Questioning  string         `json:"questioning"`
	Type         string         `json:"type"`
	Image        *string        `json:"image"`
	Alternatives pq.StringArray `json:"alternatives"`
	Answer       *string        `json:"answer"`
	Discipline   string         `json:"discipline"`
}

type UpdateQuestionInput struct {
	Questioning  string         `json:"questioning"`
	Type         string         `json:"type"`
	Image        *string        `json:"image"`
	Alternatives pq.StringArray `json:"alternatives"`
	Answer       *string        `json:"answer"`
	Discipline   string         `json:"discipline"`
}

type QuestionOutput struct {
	ID           string         `json:"id"`
	Questioning  string         `json:"questioning"`
	Type         string         `json:"type"`
	Image        *string        `json:"image"`
	Alternatives pq.StringArray `json:"alternatives"`
	Answer       *string        `json:"answer"`
	Discipline   string         `json:"discipline"`
}
