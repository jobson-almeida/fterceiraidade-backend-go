package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteCourse struct {
	CourseRepository repository.ICourseRepository
}

func NewDeleteCourse(repository repository.ICourseRepository) *DeleteCourse {
	return &DeleteCourse{CourseRepository: repository}
}

func (d *DeleteCourse) Execute(input dto.IDInput) error {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.CourseRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
