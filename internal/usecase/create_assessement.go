package usecase

import (
	"time"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateAssessment struct {
	AssessmentRepository repository.IAssessmentRepository
}

func NewCreateAssessment(repository repository.IAssessmentRepository) *CreateAssessment {
	return &CreateAssessment{AssessmentRepository: repository}
}

func (c *CreateAssessment) Execute(input *dto.AssessmentInput) error {
	var quiz []*entity.Quiz
	for _, r := range input.Quiz {
		quiz = append(quiz, &entity.Quiz{
			ID:    r.ID,
			Value: r.Value,
		})
	}

	startDate, _ := time.Parse(time.DateOnly, input.StartDate)
	endDate, _ := time.Parse(time.DateOnly, input.EndDate)

	assessment, err := entity.NewAssessment(
		input.Description,
		input.Courses,
		input.Classrooms,
		startDate,
		endDate,
		quiz,
	)
	if err != nil {
		return err
	}

	err = c.AssessmentRepository.Create(assessment)
	if err != nil {
		return err
	}
	return nil
}
