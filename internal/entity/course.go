package entity

import (
	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
)

type Course struct {
	Base
	Name        string `json:"name" validate:"required" gorm:"type:varchar(255);unique"`
	Description string `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Image       string `json:"image" validate:"required" gorm:"type:varchar(255)"`
}

func init() {}

func NewCourse(name string, description string, image string) (*Course, error) {
	course := &Course{
		Description: description,
		Name:        name,
		Image:       image,
	}
	course.ID = uuid.New().String()

	err := util.Validation(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func UpdateCourse(description string, name string, image string) (*Course, error) {
	course := &Course{
		Description: description,
		Name:        name,
		Image:       image,
	}

	err := util.Validation(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}
