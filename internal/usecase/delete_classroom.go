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
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.ClassroomRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
