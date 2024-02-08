package repository

import (
	"errors"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"gorm.io/gorm"
)

type IClassroomRepository interface {
	Create(classroom *entity.Classroom) error
	Select() ([]*entity.Classroom, error)
	Show(id *entity.InputID) (*entity.Classroom, error)
	Update(id *entity.InputID, classroom *entity.Classroom) error
	Delete(id *entity.InputID) error
}

type ClassroomRepository struct {
	db *gorm.DB
}

func NewClassroomRepository(db *gorm.DB) IClassroomRepository {
	return &ClassroomRepository{db: db}
}

func (c *ClassroomRepository) Create(classroom *entity.Classroom) error {
	res := c.db.Create(classroom)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (c *ClassroomRepository) Select() ([]*entity.Classroom, error) {
	var classrooms []*entity.Classroom
	res := c.db.Find(&classrooms)
	if res.Error != nil {
		return nil, res.Error
	}
	return classrooms, nil
}

func (c *ClassroomRepository) Show(input *entity.InputID) (*entity.Classroom, error) {
	classroom := &entity.Classroom{}
	res := c.db.First(&classroom, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return classroom, nil
}

func (c *ClassroomRepository) Update(input *entity.InputID, classroom *entity.Classroom) error {
	res := c.db.Model(&classroom).Omit("ID").Where("id = ?", input.ID).Updates(entity.Classroom{Name: classroom.Name, Description: classroom.Description, Course: classroom.Course})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *ClassroomRepository) Delete(input *entity.InputID) error {
	classroom := &entity.Classroom{}
	res := c.db.Delete(&classroom, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
