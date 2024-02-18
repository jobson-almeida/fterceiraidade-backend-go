package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowAssessment struct {
	AssessmentRepository repository.IAssessmentRepository
}

func NewShowAssessment(repository repository.IAssessmentRepository) *ShowAssessment {
	return &ShowAssessment{AssessmentRepository: repository}
}

func (s *ShowAssessment) Execute(input dto.IDInput) (*dto.AssessmentOutput, error) {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.AssessmentRepository.Show(id)
	if err != nil {
		return nil, err
	}

	var quiz []*dto.Quiz
	for _, q := range res.Quiz {
		quiz = append(quiz, &dto.Quiz{
			ID:    q.ID,
			Value: q.Value,
		})
	}

	output := &dto.AssessmentOutput{
		ID:         res.ID,
		Courses:    res.Courses,
		Classrooms: res.Classrooms,
		StartDate:  res.StartDate,
		EndDate:    res.EndDate,
		Quiz:       quiz,
	}
	return output, nil
}
