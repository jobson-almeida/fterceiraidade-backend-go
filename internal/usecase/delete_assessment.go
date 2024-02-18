package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteAssessment struct {
	AssessmentRepository repository.IAssessmentRepository
}

func NewDeleteAssessment(repository repository.IAssessmentRepository) *DeleteAssessment {
	return &DeleteAssessment{AssessmentRepository: repository}
}

func (d *DeleteAssessment) Execute(input dto.IDInput) error {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.AssessmentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
