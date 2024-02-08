package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type CreateStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewCreateStudent(repository repository.IStudentRepository) *CreateStudent {
	return &CreateStudent{StudentRepository: repository}
}

func (c *CreateStudent) Execute(input dto.StudentInput) error {
	address := entity.DetailsAddress{
		City: input.Address.City, State: input.Address.State, Street: input.Address.Street,
	}
	student, err := entity.NewStudent(
		input.Avatar,
		input.Firstname,
		input.Lastname,
		input.Email,
		input.Phone,
		address,
	)

	if err != nil {
		return err
	}

	err = c.StudentRepository.Create(student)
	if err != nil {
		return err
	}

	return nil
}
