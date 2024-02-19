package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type UpdateTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewUpdateTeacher(repository repository.ITeacherRepository) *UpdateTeacher {
	return &UpdateTeacher{TeacherRepository: repository}
}

func (u *UpdateTeacher) Execute(where dto.IDInput, data dto.UpdateTeacherInput) error {
	id, err := entity.NewInputID(where.ID)
	if err != nil {
		return err
	}

	address := entity.DetailsAddress{
		City: data.Address.City, State: data.Address.State, Street: data.Address.Street,
	}
	teacher, err := entity.UpdateTeacher(
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

	err = u.TeacherRepository.Update(id, teacher)
	if err != nil {
		return err
	}

	return nil
}
