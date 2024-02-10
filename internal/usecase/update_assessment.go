package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateAssessment struct {
	AssessmentRepository repository.IAssessmentRepository
}

func NewUpdateAssessment(repository repository.IAssessmentRepository) *UpdateAssessment {
	return &UpdateAssessment{AssessmentRepository: repository}
}

func (u *UpdateAssessment) Execute(where dto.IDInput, data dto.UpdateAssessmentInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

	var quiz []*entity.Quiz
	for _, r := range data.Quiz {
		quiz = append(quiz, &entity.Quiz{
			ID:    r.ID,
			Value: r.Value,
		})
	}

	assessment, err := entity.UpdateAssessment(
		data.Description,
		data.Courses,
		data.Classrooms,
		data.StartDate,
		data.EndDate,
		quiz,
	)

	if err != nil {
		return err
	}

	err = u.AssessmentRepository.Update(input, assessment)
	if err != nil {
		return err
	}
	return nil
}
