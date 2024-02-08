package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"
)

type UpdateTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewUpdateTeacher(repository repository.ITeacherRepository) *UpdateTeacher {
	return &UpdateTeacher{TeacherRepository: repository}
}

func (u *UpdateTeacher) Execute(where dto.IDInput, data dto.UpdateTeacherInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

	teacher := entity.UpdateTeacher()
	teacher.Avatar = data.Avatar
	teacher.Firstname = data.Firstname
	teacher.Lastname = data.Lastname
	teacher.Email = data.Email
	teacher.Phone = data.Phone
	teacher.Address = entity.DetailsAddress{City: data.Address.City, State: data.Address.State, Street: data.Address.Street}

	err := u.TeacherRepository.Update(input, teacher)
	if err != nil {
		return err
	}
	return nil
}
