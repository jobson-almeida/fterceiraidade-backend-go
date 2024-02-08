package entity

import (
	"github.com/google/uuid"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
)

type Student struct {
	Base
	Avatar    string         `json:"avatar" validate:"" gorm:"type:varchar(255)"`
	Firstname string         `json:"firstname" validate:"required,min=4" gorm:"type:varchar(255)"`
	Lastname  string         `json:"lastname" validate:"required,min=4" gorm:"type:varchar(255)"`
	Email     string         `json:"email" validate:"required,email" gorm:"type:varchar(255);unique"`
	Phone     string         `json:"phone" validate:"required,phone" gorm:"type:varchar(255)"`
	Address   DetailsAddress `json:"address" validate:"required" gorm:"type:bytes;serializer:gob"`
}

func init() {}

func NewStudent(avatar string, firstname string, lastname string, email string, phone string, address DetailsAddress) (*Student, error) {
	student := &Student{
		Avatar:    avatar,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		Address: DetailsAddress{
			City:   address.City,
			State:  address.State,
			Street: address.Street,
		},
	}
	student.ID = uuid.New().String()

	err := util.Validation(student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func UpdateStudent(avatar string, firstname string, lastname string, email string, phone string, address DetailsAddress) (*Student, error) {
	student := &Student{
		Avatar:    avatar,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		Address: DetailsAddress{
			City:   address.City,
			State:  address.State,
			Street: address.Street,
		},
	}

	err := util.Validation(student)
	if err != nil {
		return nil, err
	}

	return student, nil
}
