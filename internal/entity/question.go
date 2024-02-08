package entity

import (
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

func NewQuestion() *Question {
	return &Question{}
}

func UpdateQuestion() *Question {
	return &Question{}
}
