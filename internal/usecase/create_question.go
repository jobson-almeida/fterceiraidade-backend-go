package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewCreateQuestion(repository repository.IQuestionRepository) *CreateQuestion {
	return &CreateQuestion{QuestionRepository: repository}
}

func (c *CreateQuestion) Execute(input dto.QuestionInput) error {
	question, err := entity.NewQuestion(
		input.Questioning,
		input.Type,
		input.Image,
		input.Alternatives,
		input.Answer,
		input.Discipline,
	)
	if err != nil {
		return err
	}

	err = c.QuestionRepository.Create(question)
	if err != nil {
		return err
	}
	return nil
}
