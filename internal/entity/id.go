package entity

import "github.com/jobson-almeida/fterceiraidade-backend-go/util"

type InputID struct {
	ID string `json:"id" validate:"uuid"`
}

func NewInputID(id string) (*InputID, error) {
	inputID := &InputID{
		ID: id,
	}

	err := util.Validation(inputID)
	if err != nil {
		return nil, err
	}

	return inputID, nil
}
