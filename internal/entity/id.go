package entity

type InputID struct {
	ID string `json:"id" validate:"required"`
}

func NewInputID() *InputID {
	return &InputID{}
}
