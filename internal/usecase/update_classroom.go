package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewUpdateClassroom(repository repository.IClassroomRepository) *UpdateClassroom {
	return &UpdateClassroom{ClassroomRepository: repository}
}

func (u *UpdateClassroom) Execute(where dto.IDInput, data dto.UpdateClassroomInput) error {
	id, err := entity.NewInputID(where.ID)
	if err != nil {
		return err
	}

	classroom, err := entity.UpdateClassroom(
		data.Name,
		data.Description,
		data.Course,
	)

	if err != nil {
		return err
	}

	err = u.ClassroomRepository.Update(id, classroom)
	if err != nil {
		return err
	}
	return nil
}
