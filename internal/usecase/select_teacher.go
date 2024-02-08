package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/repository"
)

type SelectTeachers struct {
	TeacherRepository repository.ITeacherRepository
}

func NewSelectTeachers(repository repository.ITeacherRepository) *SelectTeachers {
	return &SelectTeachers{TeacherRepository: repository}
}

func (s *SelectTeachers) Execute() ([]*dto.TeacherOutput, error) {
	res, err := s.TeacherRepository.Select()
	if err != nil {
		return []*dto.TeacherOutput{}, err
	}

	var output []*dto.TeacherOutput
	for _, r := range res {
		output = append(output, &dto.TeacherOutput{
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
