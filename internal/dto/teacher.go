package dto

type TeacherInput struct {
	ID        string         `json:"id"`
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}

type UpdateTeacherInput struct {
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}

type TeacherOutput struct {
	ID        string         `json:"id"`
	Avatar    string         `json:"avatar"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   DetailsAddress `json:"address"`
}
