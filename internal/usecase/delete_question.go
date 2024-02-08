package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"
)

type DeleteQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewDeleteQuestion(repository repository.IQuestionRepository) *DeleteQuestion {
	return &DeleteQuestion{QuestionRepository: repository}
}

func (d *DeleteQuestion) Execute(input dto.IDInput) error {
	course := entity.NewInputID()
	course.ID = input.ID
	err := d.QuestionRepository.Delete(course)
	if err != nil {
		return err
	}
	return nil
}
