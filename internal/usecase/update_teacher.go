package usecase

import (
	"fmt"

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
		fmt.Println("usecase")
		return err
	}

	/*if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
		fmt.Println("pq error:", err.Column)
		fmt.Println("pq error:", err.Constraint)
		fmt.Println("pq error:", err.DataTypeName)
		fmt.Println("pq error:", err.Detail)
		fmt.Println("pq error:", err.File)
		fmt.Println("pq error:", err.Hint)
		fmt.Println("pq error:", err.InternalPosition)
		fmt.Println("pq error:", err.InternalQuery)
		fmt.Println("pq error:", err.Line)
		fmt.Println("pq error:", err.Message)
		fmt.Println("pq error:", err.Position)
		fmt.Println("pq error:", err.Routine)
		fmt.Println("pq error:", err.Schema)
		fmt.Println("pq error:", err.Severity)
		fmt.Println("pq error:", err.Table)
		fmt.Println("pq error:", err.Where)

		fmt.Println("usecase")
		return err
	}*/

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
