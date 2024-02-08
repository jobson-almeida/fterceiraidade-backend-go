package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewCreateTeacher(repository repository.ITeacherRepository) *CreateTeacher {
	return &CreateTeacher{TeacherRepository: repository}
}

func (c *CreateTeacher) Execute(input dto.TeacherInput) error {
	address := entity.DetailsAddress{
		City: input.Address.City, State: input.Address.State, Street: input.Address.Street,
	}
	teacher, err := entity.NewTeacher(
		input.Avatar,
		input.Firstname,
		input.Lastname,
		input.Email,
		input.Phone,
		address,
	)

	if err != nil {
		return err
	}

	err = c.TeacherRepository.Create(teacher)
	if err != nil {
		return err
	}

	return nil
}
