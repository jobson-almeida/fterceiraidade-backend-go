package dto

import (
	"github.com/lib/pq"
)

type AssessmentInput struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"startdate"`
	EndDate     string         `json:"enddate"`
	Quiz        []*Quiz        `json:"quiz"`
}

type UpdateAssessmentInput struct {
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"startdate"`
	EndDate     string         `json:"enddate"`
	Quiz        []*Quiz        `json:"quiz"`
}

type AssessmentOutput struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Courses     pq.StringArray `json:"courses"`
	Classrooms  pq.StringArray `json:"classrooms"`
	StartDate   string         `json:"startdate"`
	EndDate     string         `json:"enddate"`
	Quiz        []*Quiz        `json:"quiz"`
}
