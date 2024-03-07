package repository

import (
	"errors"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"gorm.io/gorm"
)

type ICourseRepository interface {
	Create(course *entity.Course) error
	Select() ([]*entity.Course, error)
	Show(id *entity.InputID) (*entity.Course, error)
	Update(id *entity.InputID, course *entity.Course) error
	Delete(id *entity.InputID) error
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (c *CourseRepository) Create(course *entity.Course) error {
	res := c.db.Create(course)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (c *CourseRepository) Select() ([]*entity.Course, error) {
	var courses []*entity.Course
	res := c.db.Find(&courses)
	if res.Error != nil {
		return nil, res.Error
	}
	return courses, nil
}

func (c *CourseRepository) Show(input *entity.InputID) (*entity.Course, error) {
	course := &entity.Course{}
	res := c.db.First(&course, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return course, nil
}

func (c *CourseRepository) Update(input *entity.InputID, course *entity.Course) error {
	res := c.db.Model(&course).Omit("ID").Where("id = ?", input.ID).Updates(entity.Course{Name: course.Name, Description: course.Description, Image: course.Image})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CourseRepository) Delete(input *entity.InputID) error {
	course := &entity.Course{}

	res := c.db.First(&course, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}

	res = c.db.Delete(&course, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
