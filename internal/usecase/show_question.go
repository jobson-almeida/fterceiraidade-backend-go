package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowQuestion struct {
	QuestionRepository repository.IQuestionRepository
}

func NewShowQuestion(repository repository.IQuestionRepository) *ShowQuestion {
	return &ShowQuestion{QuestionRepository: repository}
}

func (s *ShowQuestion) Execute(input dto.IDInput) (*dto.QuestionOutput, error) {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.QuestionRepository.Show(id)
	if err != nil {
		return nil, err
	}

	output := &dto.QuestionOutput{
		ID:           res.ID,
		Questioning:  res.Questioning,
		Type:         res.Type,
		Image:        res.Image,
		Alternatives: res.Alternatives,
		Answer:       res.Answer,
		Discipline:   res.Discipline,
	}
	return output, nil
}
