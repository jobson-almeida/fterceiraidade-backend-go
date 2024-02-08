package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type SelectCourses struct {
	CourseRepository repository.ICourseRepository
}

func NewSelectCourses(repository repository.ICourseRepository) *SelectCourses {
	return &SelectCourses{CourseRepository: repository}
}

func (s *SelectCourses) Execute() ([]*dto.CourseOutput, error) {
	res, err := s.CourseRepository.Select()
	if err != nil {
		return []*dto.CourseOutput{}, err
	}

	var output []*dto.CourseOutput
	for _, r := range res {
		output = append(output, &dto.CourseOutput{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Image:       r.Image,
		})
	}
	return output, nil
}
