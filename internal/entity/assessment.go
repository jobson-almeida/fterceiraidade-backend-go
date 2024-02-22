package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
	"github.com/lib/pq"
)

type Assessment struct {
	Base
	Description string         `json:"description" validate:"required,max=22" gorm:"type:varchar(22)"`
	Courses     pq.StringArray `json:"courses" validate:"required" gorm:"type:text[]"`
	Classrooms  pq.StringArray `json:"classrooms" validate:"required" gorm:"type:text[]"`
	StartDate   time.Time      `json:"start_date" validate:"required" gorm:"type:date"`
	EndDate     time.Time      `json:"end_date" validate:"required,gtefield=StartDate" gorm:"type:date"`
	Quiz        []*Quiz        `json:"quiz" gorm:"serializer:json"`
}

func init() {}

func NewAssessment(description string, courses pq.StringArray, classrooms pq.StringArray, startDate time.Time, endDate time.Time, quiz []*Quiz) (*Assessment, error) {
	var questions []*Quiz
	for _, r := range quiz {
		questions = append(questions, &Quiz{
			ID:    r.ID,
			Value: r.Value,
		})
	}

	assessment := &Assessment{
		Description: description,
		Courses:     courses,
		Classrooms:  classrooms,
		StartDate:   startDate,
		EndDate:     endDate,
		Quiz:        questions,
	}
	assessment.ID = uuid.New().String()

	err := util.Validation(assessment)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}

func UpdateAssessment(description string, courses pq.StringArray, classrooms pq.StringArray, startDate time.Time, endDate time.Time, quiz []*Quiz) (*Assessment, error) {
	var questions []*Quiz
	for _, r := range quiz {
		questions = append(questions, &Quiz{
			ID:    r.ID,
			Value: r.Value,
		})
	}

	assessment := &Assessment{
		Description: description,
		Courses:     courses,
		Classrooms:  classrooms,
		StartDate:   startDate,
		EndDate:     endDate,
		Quiz:        questions,
	}

	err := util.Validation(assessment)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}
