package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewShowTeacher(repository repository.ITeacherRepository) *ShowTeacher {
	return &ShowTeacher{TeacherRepository: repository}
}

func (s *ShowTeacher) Execute(input dto.IDInput) (*dto.TeacherOutput, error) {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.TeacherRepository.Show(id)
	if err != nil {
		return nil, err
	}

	output := &dto.TeacherOutput{
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
