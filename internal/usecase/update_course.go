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
	id, err := entity.NewInputID(where.ID)
	if err != nil {
		return err
	}

	course, err := entity.UpdateCourse(
		data.Name,
		data.Description,
		data.Image,
	)

	if err != nil {
		return err
	}

	err = u.CourseRepository.Update(id, course)
	if err != nil {
		return err
	}
	return nil
}
