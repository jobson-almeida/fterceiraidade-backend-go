package dto

type ClassroomInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Course      string `json:"course"`
}

type UpdateClassroomInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Course      string `json:"course"`
}

type ClassroomOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Course      string `json:"course"`
}
