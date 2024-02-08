package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	entity "github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowCourse struct {
	CourseRepository repository.ICourseRepository
}

func NewShowCourse(repository repository.ICourseRepository) *ShowCourse {
	return &ShowCourse{CourseRepository: repository}
}

func (s *ShowCourse) Execute(input dto.IDInput) (*dto.CourseOutput, error) {
	course := entity.NewInputID()
	course.ID = input.ID

	res, err := s.CourseRepository.Show(course)
	if err != nil {
		return nil, err
	}

	output := &dto.CourseOutput{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Image:       res.Image,
	}
	return output, nil
}
