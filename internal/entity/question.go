package entity

import (
	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
	"github.com/lib/pq"
)

type Question struct {
	Base
	Questioning  string         `json:"questioning" validate:"required" gorm:"type:varchar(255)"`
	Type         string         `json:"type" validate:"required" gorm:"type:varchar(255)"`
	Image        *string        `json:"image" gorm:"type:varchar(255)"`
	Alternatives pq.StringArray `json:"alternatives" gorm:"type:text[]"`
	Answer       *string        `json:"answer" gorm:"type:varchar(255)"`
	Discipline   string         `json:"discipline" validate:"required" gorm:"type:varchar(255)"`
}

func init() {}

func NewQuestion(questioning string, type_ string, image *string, alternatives pq.StringArray, answer *string, discipline string) (*Question, error) {
	question := &Question{
		Questioning:  questioning,
		Type:         type_,
		Image:        image,
		Alternatives: alternatives,
		Answer:       answer,
		Discipline:   discipline,
	}
	question.ID = uuid.New().String()

	err := util.Validation(question)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func UpdateQuestion(questioning string, type_ string, image *string, alternatives pq.StringArray, answer *string, discipline string) (*Question, error) {
	question := &Question{
		Questioning:  questioning,
		Type:         type_,
		Image:        image,
		Alternatives: alternatives,
		Answer:       answer,
		Discipline:   discipline,
	}

	err := util.Validation(question)
	if err != nil {
		return nil, err
	}

	return question, nil
}
