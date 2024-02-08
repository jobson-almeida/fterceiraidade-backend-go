package entity

type Course struct {
	Base
	Name        string `json:"name" validate:"required" gorm:"type:varchar(255);unique"`
	Description string `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Image       string `json:"image" validate:"required" gorm:"type:varchar(255)"`
}

func init() {}

func NewCourse() *Course {
	return &Course{}
}

func UpdateCourse() *Course {
	return &Course{}
}
