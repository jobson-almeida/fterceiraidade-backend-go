package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewDeleteQuestion(repository repository.IQuestionRepository) *DeleteQuestion {
	return &DeleteQuestion{QuestionRepository: repository}
}

func (d *DeleteQuestion) Execute(input dto.IDInput) error {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.QuestionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
