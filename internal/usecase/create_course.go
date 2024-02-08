package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"

	"github.com/google/uuid"
)

type CreateCourse struct {
	CourseRepository repository.ICourseRepository
}

func NewCreateCourse(repository repository.ICourseRepository) *CreateCourse {
	return &CreateCourse{CourseRepository: repository}
}

func (c *CreateCourse) Execute(input dto.CourseInput) error {
	course := entity.NewCourse()
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Image = input.Image

	err := c.CourseRepository.Create(course)
	if err != nil {
		return err
	}
	return nil
}
