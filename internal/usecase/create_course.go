package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateCourse struct {
	CourseRepository repository.ICourseRepository
}

func NewCreateCourse(repository repository.ICourseRepository) *CreateCourse {
	return &CreateCourse{CourseRepository: repository}
}

func (c *CreateCourse) Execute(input dto.CourseInput) error {
	course, err := entity.NewCourse(
		input.Name,
		input.Description,
		input.Image,
	)
	if err != nil {
		return err
	}

	err = c.CourseRepository.Create(course)
	if err != nil {
		return err
	}
	return nil
}
