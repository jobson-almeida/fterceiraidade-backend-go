package repository

import (
	"errors"
	"fterceiraidade-backend-go/internal/entity"

	"gorm.io/gorm"
)

type IStudentRepository interface {
	Create(student *entity.Student) error
	Select() ([]*entity.Student, error)
	Show(id *entity.InputID) (*entity.Student, error)
	Update(id *entity.InputID, student *entity.Student) error
	Delete(id *entity.InputID) error
}

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (s *StudentRepository) Create(student *entity.Student) error {
	res := s.db.Create(student)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (s *StudentRepository) Select() ([]*entity.Student, error) {
	var students []*entity.Student
	res := s.db.Find(&students)
	if res.Error != nil {
		return nil, res.Error
	}
	return students, nil
}

func (s *StudentRepository) Show(input *entity.InputID) (*entity.Student, error) {
	student := &entity.Student{}
	res := s.db.First(&student, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return student, nil
}

func (s *StudentRepository) Update(input *entity.InputID, student *entity.Student) error {
	res := s.db.Model(&student).Omit("ID").Where("id = ?", input.ID).Updates(entity.Student{
		Avatar:    student.Avatar,
		Firstname: student.Firstname,
		Lastname:  student.Lastname,
		Email:     student.Email,
		Phone:     student.Phone,
		Address: entity.DetailsAddress{
			City:   student.Address.City,
			State:  student.Address.State,
			Street: student.Address.Street,
		},
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *StudentRepository) Delete(input *entity.InputID) error {
	student := &entity.Student{}
	res := s.db.Delete(&student, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
