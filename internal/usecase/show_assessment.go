package usecase

import (
	"time"

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

	startDate := time.Time(res.StartDate).Format(time.DateOnly)
	endDate := time.Time(res.EndDate).Format(time.DateOnly)

	output := &dto.AssessmentOutput{
		ID:         res.ID,
		Courses:    res.Courses,
		Classrooms: res.Classrooms,
		StartDate:  startDate,
		EndDate:    endDate,
		Quiz:       quiz,
	}
	return output, nil
}
