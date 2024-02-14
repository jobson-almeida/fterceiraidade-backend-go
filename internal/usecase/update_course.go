package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateCourse struct {
	CourseRepository repository.ICourseRepository
}

func NewUpdateCourse(repository repository.ICourseRepository) *UpdateCourse {
	return &UpdateCourse{CourseRepository: repository}
}

func (u *UpdateCourse) Execute(where dto.IDInput, data dto.UpdateCourseInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

	course, err := entity.NewCourse(
		data.Name,
		data.Description,
		data.Image,
	)
	if err != nil {
		return err
	}

	err = u.CourseRepository.Update(input, course)
	if err != nil {
		return err
	}
	return nil
}
