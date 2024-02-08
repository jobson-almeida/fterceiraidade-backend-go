package usecase

import (
	"fterceiraidade-backend-go/internal/dto"
	"fterceiraidade-backend-go/internal/entity"
	"fterceiraidade-backend-go/internal/repository"
)

type UpdateStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewUpdateStudent(repository repository.IStudentRepository) *UpdateStudent {
	return &UpdateStudent{StudentRepository: repository}
}

func (u *UpdateStudent) Execute(where dto.IDInput, data dto.UpdateStudentInput) error {
	input := entity.NewInputID()
	input.ID = where.ID

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

	err = u.StudentRepository.Update(input, student)
	if err != nil {
		return err
	}

	return nil
}

/*input := entity.NewInputID()
	input.ID = where.ID

	student := entity.UpdateStudent()
	student.Avatar = data.Avatar
	student.Firstname = data.Firstname
	student.Lastname = data.Lastname
	student.Email = data.Email
	student.Phone = data.Phone
	student.Address = entity.DetailsAddress{City: data.Address.City, State: data.Address.State, Street: data.Address.Street}

	err := cv.CustomValidator(student)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = u.StudentRepository.Update(input, student)
	if err != nil {
		return err
	}
	return nil
}
*/
