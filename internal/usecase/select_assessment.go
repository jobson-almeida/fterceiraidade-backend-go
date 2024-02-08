package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type SelectAssessments struct {
	AssessmentRepository repository.IAssessmentRepository
}

func NewSelectAssessments(repository repository.IAssessmentRepository) *SelectAssessments {
	return &SelectAssessments{AssessmentRepository: repository}
}

func (s *SelectAssessments) Execute() ([]*dto.AssessmentOutput, error) {
	res, err := s.AssessmentRepository.Select()
	if err != nil {
		return []*dto.AssessmentOutput{}, err
	}

	var output []*dto.AssessmentOutput

	for _, r := range res {
		var quiz []*dto.Quiz
		for _, q := range r.Quiz {
			quiz = append(quiz, &dto.Quiz{
				ID:    q.ID,
				Value: q.Value,
			})
		}

		output = append(output, &dto.AssessmentOutput{
			ID:          r.ID,
			Description: r.Description,
			Courses:     r.Courses,
			Classrooms:  r.Classrooms,
			StartDate:   r.StartDate,
			EndDate:     r.EndDate,
			Quiz:        quiz,
		})

	}
	return output, nil
}
