package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	entity "github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewUpdateClassroom(repository repository.IClassroomRepository) *UpdateClassroom {
	return &UpdateClassroom{ClassroomRepository: repository}
}

func (u *UpdateClassroom) Execute(where dto.IDInput, data dto.UpdateClassroomInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

	course := entity.UpdateClassroom()
	course.Name = data.Name
	course.Description = data.Description
	course.Course = data.Course

	err := u.ClassroomRepository.Update(input, course)
	if err != nil {
		return err
	}
	return nil
}
