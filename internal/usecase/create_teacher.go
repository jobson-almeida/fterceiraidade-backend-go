package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"

	"github.com/google/uuid"
)

type CreateTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewCreateTeacher(repository repository.ITeacherRepository) *CreateTeacher {
	return &CreateTeacher{TeacherRepository: repository}
}

func (c *CreateTeacher) Execute(input dto.TeacherInput) error {
	teacher := entity.NewTeacher()
	teacher.ID = uuid.New().String()
	teacher.Avatar = input.Avatar
	teacher.Firstname = input.Firstname
	teacher.Lastname = input.Lastname
	teacher.Email = input.Email
	teacher.Phone = input.Phone
	teacher.Address = entity.DetailsAddress{City: input.Address.City, State: input.Address.State, Street: input.Address.Street}

	err := c.TeacherRepository.Create(teacher)
	if err != nil {
		return err
	}
	return nil
}
