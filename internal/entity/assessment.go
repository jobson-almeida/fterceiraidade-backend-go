package entity

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
	"github.com/lib/pq"
)

type Assessment struct {
	Base
	Description string         `json:"description" validate:"required,max=22" gorm:"type:varchar(22)"`
	Courses     pq.StringArray `json:"courses" validate:"required,array_uuid" gorm:"type:text[]"`
	Classrooms  pq.StringArray `json:"classrooms" validate:"required,array_uuid" gorm:"type:text[]"`
	StartDate   string         `json:"startdate" validate:"required" gorm:"type:varchar"`
	EndDate     string         `json:"enddate" validate:"required" gorm:"type:varchar"`
	Quiz        []*Quiz        `json:"quiz" gorm:"serializer:json"`
}

func init() {}

func NewAssessment(description string, courses pq.StringArray, classrooms pq.StringArray, startDate string, endDate string, quiz []*Quiz) (*Assessment, error) {
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
		fmt.Println(err)
		return nil, err
	}

	return assessment, nil
}

func UpdateAssessment(description string, courses pq.StringArray, classrooms pq.StringArray, startDate string, endDate string, quiz []*Quiz) (*Assessment, error) {
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
		fmt.Println(err)
		return nil, err
	}

	return assessment, nil
}
