package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
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

	question, err := entity.UpdateQuestion(
		data.Questioning,
		data.Type,
		data.Image,
		data.Alternatives,
		data.Answer,
		data.Discipline,
	)
	if err != nil {
		return err
	}

	err = u.QuestionRepository.Update(input, question)
	if err != nil {
		return err
	}
	return nil
}
