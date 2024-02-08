package usecase

import (
	dto "github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type SelectStudents struct {
	StudentRepository repository.IStudentRepository
}

func NewSelectStudents(repository repository.IStudentRepository) *SelectStudents {
	return &SelectStudents{StudentRepository: repository}
}

func (s *SelectStudents) Execute() ([]*dto.StudentOutput, error) {
	res, err := s.StudentRepository.Select()
	if err != nil {
		return []*dto.StudentOutput{}, err
	}

	var output []*dto.StudentOutput
	for _, r := range res {
		output = append(output, &dto.StudentOutput{
			ID:        r.ID,
			Avatar:    r.Avatar,
			Firstname: r.Firstname,
			Lastname:  r.Lastname,
			Email:     r.Email,
			Phone:     r.Phone,
			Address:   dto.DetailsAddress{City: r.Address.City, State: r.Address.State, Street: r.Address.Street},
		})
	}
	return output, nil
}
