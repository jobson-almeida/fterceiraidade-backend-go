package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	entity "github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewUpdateQuestion(repository repository.IQuestionRepository) *UpdateQuestion {
	return &UpdateQuestion{QuestionRepository: repository}
}

func (u *UpdateQuestion) Execute(where dto.IDInput, data dto.UpdateQuestionInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

	question := entity.UpdateQuestion()
	question.Questioning = data.Questioning
	question.Type = data.Type
	question.Image = data.Image
	question.Alternatives = data.Alternatives
	question.Answer = data.Answer
	question.Discipline = data.Discipline

	err := u.QuestionRepository.Update(input, question)
	if err != nil {
		return err
	}
	return nil
}
