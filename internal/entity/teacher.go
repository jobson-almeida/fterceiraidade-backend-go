package entity

type Teacher struct {
	Base
	Avatar    string         `json:"avatar" validate:"required" gorm:"type:varchar(255)"`
	Firstname string         `json:"firstname" validate:"required" gorm:"type:varchar(255)"`
	Lastname  string         `json:"lastname" validate:"required" gorm:"type:varchar(255)"`
	Email     string         `json:"email" validate:"required" gorm:"type:varchar(255);unique"`
	Phone     string         `json:"phone" validate:"required" gorm:"type:varchar(255)"`
	Address   DetailsAddress `json:"address" validate:"required" gorm:"type:bytes;serializer:gob"`
}

func init() {}

func NewTeacher() *Teacher {
	return &Teacher{}
}

func UpdateTeacher() *Teacher {
	return &Teacher{}
}
