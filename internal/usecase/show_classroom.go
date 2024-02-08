package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"
)

type ShowClassroom struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewShowClassroom(repository repository.IClassroomRepository) *ShowClassroom {
	return &ShowClassroom{ClassroomRepository: repository}
}

func (s *ShowClassroom) Execute(input dto.IDInput) (*dto.ClassroomOutput, error) {
	course := entity.NewInputID()
	course.ID = input.ID

	res, err := s.ClassroomRepository.Show(course)
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
