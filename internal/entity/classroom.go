package entity

import (
	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
)

type Classroom struct {
	Base
	Name        string `json:"name" validate:"required" gorm:"type:varchar(255);unique"`
	Description string `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Course      string `json:"course" validate:"required,uuid" gorm:"type:varchar(255)"`
}

func init() {}

func NewClassroom(name string, description string, course string) (*Classroom, error) {
	classname := &Classroom{
		Name:        name,
		Description: description,
		Course:      course,
	}
	classname.ID = uuid.New().String()

	err := util.Validation(classname)
	if err != nil {
		return nil, err
	}

	return classname, nil
}

func UpdateClassroom(name string, description string, course string) (*Classroom, error) {
	classname := &Classroom{
		Name:        name,
		Description: description,
		Course:      course,
	}

	err := util.Validation(classname)
	if err != nil {
		return nil, err
	}

	return classname, nil
}
