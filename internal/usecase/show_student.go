package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewShowStudent(repository repository.IStudentRepository) *ShowStudent {
	return &ShowStudent{StudentRepository: repository}
}

func (s *ShowStudent) Execute(input dto.IDInput) (*dto.StudentOutput, error) {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.StudentRepository.Show(id)
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
