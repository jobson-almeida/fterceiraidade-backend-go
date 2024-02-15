package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewCreateClassroom(repository repository.IClassroomRepository) *CreateClassroom {
	return &CreateClassroom{ClassroomRepository: repository}
}

func (c *CreateClassroom) Execute(input dto.ClassroomInput) error {
	classroom, err := entity.NewClassroom(
		input.Name,
		input.Description,
		input.Course,
	)
	if err != nil {
		return err
	}

	err = c.ClassroomRepository.Create(classroom)
	if err != nil {
		return err
	}
	return nil
}
