package entity

import (
	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
	"github.com/lib/pq"
)

type Assessment struct {
	Base
	Description string         `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Courses     pq.StringArray `json:"image" validate:"required" gorm:"type:text[]"`
	Classrooms  pq.StringArray `json:"classrooms" validate:"required" gorm:"type:text[]"`
	StartDate   string         `json:"startdate" validate:"required" gorm:"type:varchar"`
	EndDate     string         `json:"enddate" validate:"required" gorm:"type:varchar"`
	Quiz        []*Quiz        `json:"quiz" gorm:"serializer:json"`
}

func init() {}

/*
func cNewAssessment() *Assessment {
	return &Assessment{}
}*/

func UpdateAssessment() *Assessment {
	return &Assessment{}
}

func NewAssessment(description string, courses pq.StringArray, classrooms pq.StringArray, startDate string, endDate string, quiz []*Quiz) (*Assessment, error) {
	assessment := &Assessment{
		Description: description,
		Courses:     courses,
		Classrooms:  classrooms,
		StartDate:   startDate,
		EndDate:     endDate,
		Quiz:        quiz,
	}
	assessment.ID = uuid.New().String()

	err := util.Validation(assessment)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}
