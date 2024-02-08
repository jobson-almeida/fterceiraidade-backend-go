package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type SelectQuestions struct {
	QuestionRepository repository.IQuestionRepository
}

func NewSelectQuestions(repository repository.IQuestionRepository) *SelectQuestions {
	return &SelectQuestions{QuestionRepository: repository}
}

func (s *SelectQuestions) Execute() ([]*dto.QuestionOutput, error) {
	res, err := s.QuestionRepository.Select()
	if err != nil {
		return []*dto.QuestionOutput{}, err
	}

	var output []*dto.QuestionOutput
	for _, r := range res {
		output = append(output, &dto.QuestionOutput{
			ID:           r.ID,
			Questioning:  r.Questioning,
			Type:         r.Type,
			Image:        r.Image,
			Alternatives: r.Alternatives,
			Answer:       r.Answer,
			Discipline:   r.Discipline,
		})
	}
	return output, nil
}
