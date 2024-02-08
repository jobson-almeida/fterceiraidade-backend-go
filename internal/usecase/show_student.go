package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	entity "github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewShowStudent(repository repository.IStudentRepository) *ShowStudent {
	return &ShowStudent{StudentRepository: repository}
}

func (s *ShowStudent) Execute(input dto.IDInput) (*dto.StudentOutput, error) {
	teacher := entity.NewInputID()
	teacher.ID = input.ID

	res, err := s.StudentRepository.Show(teacher)
	if err != nil {
		return nil, err
	}

	output := &dto.StudentOutput{
		ID:        res.ID,
		Avatar:    res.Avatar,
		Firstname: res.Firstname,
		Lastname:  res.Lastname,
		Email:     res.Email,
		Phone:     res.Phone,
		Address:   dto.DetailsAddress{City: res.Address.City, State: res.Address.State, Street: res.Address.Street},
	}
	return output, nil
}
