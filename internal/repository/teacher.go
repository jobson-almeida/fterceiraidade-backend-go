package repository

import (
	"errors"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"gorm.io/gorm"
)

type ITeacherRepository interface {
	Create(teacher *entity.Teacher) error
	Select() ([]*entity.Teacher, error)
	Show(id *entity.InputID) (*entity.Teacher, error)
	Update(id *entity.InputID, teacher *entity.Teacher) error
	Delete(id *entity.InputID) error
}

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (t *TeacherRepository) Create(teacher *entity.Teacher) error {
	res := t.db.Create(teacher)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (t *TeacherRepository) Select() ([]*entity.Teacher, error) {
	var teachers []*entity.Teacher
	res := t.db.Find(&teachers)
	if res.Error != nil {
		return nil, res.Error
	}
	return teachers, nil
}

func (t *TeacherRepository) Show(input *entity.InputID) (*entity.Teacher, error) {
	teacher := &entity.Teacher{}
	res := t.db.First(&teacher, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return teacher, nil
}

func (t *TeacherRepository) Update(input *entity.InputID, teacher *entity.Teacher) error {
	res := t.db.Model(&teacher).Omit("ID").Where("id = ?", input.ID).Updates(entity.Teacher{
		Avatar:    teacher.Avatar,
		Firstname: teacher.Firstname,
		Lastname:  teacher.Lastname,
		Email:     teacher.Email,
		Phone:     teacher.Phone,
		Address: entity.DetailsAddress{
			City:   teacher.Address.City,
			State:  teacher.Address.State,
			Street: teacher.Address.Street,
		},
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *TeacherRepository) Delete(input *entity.InputID) error {
	teacher := &entity.Teacher{}

	res := t.db.First(&teacher, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}

	res = t.db.Delete(&teacher, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
