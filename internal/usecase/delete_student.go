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
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.StudentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
