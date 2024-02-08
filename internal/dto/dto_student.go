package dto

type StudentInput struct {
	ID        string         `json:"id"`
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}

type UpdateStudentInput struct {
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}

type StudentOutput struct {
	ID        string         `json:"id"`
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}
