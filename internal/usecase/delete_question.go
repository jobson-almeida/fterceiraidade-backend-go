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
	question := entity.NewInputID()
	question.ID = input.ID
	err := d.QuestionRepository.Delete(question)
	if err != nil {
		return err
	}
	return nil
}
