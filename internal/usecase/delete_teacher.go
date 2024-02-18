package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type DeleteTeacher struct {
	TeacherRepository repository.ITeacherRepository
}

func NewDeleteTeacher(repository repository.ITeacherRepository) *DeleteTeacher {
	return &DeleteTeacher{TeacherRepository: repository}
}

func (d *DeleteTeacher) Execute(input dto.IDInput) error {
	id, err := entity.NewInputID(input.ID)
	if err != nil {
		return err
	}

	err = d.TeacherRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
