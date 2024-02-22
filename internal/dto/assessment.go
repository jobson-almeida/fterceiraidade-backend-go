package dto

import (
	"github.com/lib/pq"
)

type AssessmentInput struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Quiz        []*Quiz        `json:"quiz"`
}

type UpdateAssessmentInput struct {
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Quiz        []*Quiz        `json:"quiz"`
}

type AssessmentOutput struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Quiz        []*Quiz        `json:"quiz"`
}
