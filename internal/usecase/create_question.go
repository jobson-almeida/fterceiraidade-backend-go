package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"

	"github.com/google/uuid"
)

type CreateQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewCreateQuestion(repository repository.IQuestionRepository) *CreateQuestion {
	return &CreateQuestion{QuestionRepository: repository}
}

func (c *CreateQuestion) Execute(input dto.QuestionInput) error {
	question := entity.NewQuestion()
	question.ID = uuid.New().String()
	question.Questioning = input.Questioning
	question.Type = input.Type
	question.Image = input.Image
	question.Alternatives = input.Alternatives
	question.Answer = input.Answer
	question.Discipline = input.Discipline

	err := c.QuestionRepository.Create(question)
	if err != nil {
		return err
	}
	return nil
}
