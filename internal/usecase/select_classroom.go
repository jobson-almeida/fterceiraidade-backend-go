package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/repository"
)

type SelectClassrooms struct {
	ClassroomRepository repository.IClassroomRepository
}

func NewSelectClassrooms(repository repository.IClassroomRepository) *SelectClassrooms {
	return &SelectClassrooms{ClassroomRepository: repository}
}

func (s *SelectClassrooms) Execute() ([]*dto.ClassroomOutput, error) {
	res, err := s.ClassroomRepository.Select()
	if err != nil {
		return []*dto.ClassroomOutput{}, err
	}

	var output []*dto.ClassroomOutput
	for _, r := range res {
		output = append(output, &dto.ClassroomOutput{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Course:      r.Course,
		})
	}
	return output, nil
}
