package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewUpdateStudent(repository repository.IStudentRepository) *UpdateStudent {
	return &UpdateStudent{StudentRepository: repository}
}

func (u *UpdateStudent) Execute(where dto.IDInput, data dto.UpdateStudentInput) error {
	id, err := entity.NewInputID(where.ID)
	if err != nil {
		return err
	}

	address := entity.DetailsAddress{
		City: data.Address.City, State: data.Address.State, Street: data.Address.Street,
	}
	student, err := entity.UpdateStudent(
		data.Avatar,
		data.Firstname,
		data.Lastname,
		data.Email,
		data.Phone,
		address,
	)

	if err != nil {
		return err
	}

	err = u.StudentRepository.Update(id, student)
	if err != nil {
		return err
	}

	return nil
}
