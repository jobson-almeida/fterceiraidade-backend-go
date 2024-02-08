package usecase

import (
	//"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	//"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	//"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"

	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"

	"github.com/google/uuid"
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

	assessment := entity.NewAssessment()
	assessment.ID = uuid.New().String()
	assessment.Description = input.Description
	assessment.Courses = input.Courses
	assessment.Classrooms = input.Classrooms
	assessment.StartDate = input.StartDate
	assessment.EndDate = input.EndDate
	assessment.Quiz = quiz

	err := c.AssessmentRepository.Create(assessment)
	if err != nil {
		return err
	}
	return nil
}
