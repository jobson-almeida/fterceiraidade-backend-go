package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"

	"github.com/google/uuid"
)

type CreateClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewCreateClassroom(repository repository.IClassroomRepository) *CreateClassroom {
	return &CreateClassroom{ClassroomRepository: repository}
}

func (c *CreateClassroom) Execute(input dto.ClassroomInput) error {
	classroom := entity.NewClassroom()
	classroom.ID = uuid.New().String()
	classroom.Name = input.Name
	classroom.Description = input.Description
	classroom.Course = input.Course

	err := c.ClassroomRepository.Create(classroom)
	if err != nil {
		return err
	}
	return nil
}
