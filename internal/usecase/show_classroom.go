package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type ShowClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewShowClassroom(repository repository.IClassroomRepository) *ShowClassroom {
	return &ShowClassroom{ClassroomRepository: repository}
}

func (s *ShowClassroom) Execute(input dto.IDInput) (*dto.ClassroomOutput, error) {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.ClassroomRepository.Show(id)
	if err != nil {
		return nil, err
	}

	output := &dto.ClassroomOutput{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Course:      res.Course,
	}
	return output, nil
}
