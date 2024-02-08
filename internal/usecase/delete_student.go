package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteStudent struct {
	StudentRepository repository.IStudentRepository
}

func NewDeleteStudent(repository repository.IStudentRepository) *DeleteStudent {
	return &DeleteStudent{StudentRepository: repository}
}

func (d *DeleteStudent) Execute(input dto.IDInput) error {
	teacher := entity.NewInputID()
	teacher.ID = input.ID
	err := d.StudentRepository.Delete(teacher)
	if err != nil {
		return err
	}
	return nil
}
