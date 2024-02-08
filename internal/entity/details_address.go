package entity

type DetailsAddress struct {
	City   string `json:"city" validate:"required" gorm:"type:varchar(255)"`
	State  string `json:"state" validate:"required" gorm:"type:varchar(255)"`
	Street string `json:"street" validate:"required" gorm:"type:varchar(255)"`
}
