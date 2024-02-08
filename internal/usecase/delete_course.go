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
	course := entity.NewInputID()
	course.ID = input.ID
	err := d.CourseRepository.Delete(course)
	if err != nil {
		return err
	}
	return nil
}