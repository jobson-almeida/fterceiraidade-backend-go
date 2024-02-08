package entity

type InputID struct {
	ID string `json:"id" validate:"required"`
}

func NewInputID() *InputID {
	return &InputID{}
}

/*
func NewInputID() (*InputID, *validator.ErrResponse) {
	uuid, _ := regexp.MatchString("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", input.ID)
	if !uuid {
		return validator.ToErrResponse(err), nil
	}
	return &InputID{}, nil
}
*/
