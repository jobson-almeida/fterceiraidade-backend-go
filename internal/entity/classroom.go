package entity

type Classroom struct {
	Base
	Name        string `json:"name" validate:"required" gorm:"type:varchar(255);unique"`
	Description string `json:"description" validate:"required" gorm:"type:varchar(255)"`
	Course      string `json:"course" validate:"required" gorm:"type:varchar(255)"`
}

func init() {}

func NewClassroom() *Classroom {
	return &Classroom{}
}

func UpdateClassroom() *Classroom {
	return &Classroom{}
}
