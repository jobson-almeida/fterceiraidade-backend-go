package entity

type Quiz struct {
	ID    string `json:"id" validate:"required" gorm:"type:uuid"`
	Value int    `json:"value" validate:"required" gorm:"type:int"`
}

func NewQuiz() []Quiz {
	return []Quiz{}
}
