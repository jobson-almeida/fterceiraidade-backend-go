package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewDeleteClassroom(repository repository.IClassroomRepository) *DeleteClassroom {
	return &DeleteClassroom{ClassroomRepository: repository}
}

func (d *DeleteClassroom) Execute(input dto.IDInput) error {
	course := entity.NewInputID()
	course.ID = input.ID
	err := d.ClassroomRepository.Delete(course)
	if err != nil {
		return err
	}
	return nil
}
