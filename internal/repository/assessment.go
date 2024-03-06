package repository

import (
	"errors"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"

	"gorm.io/gorm"
)

type IAssessmentRepository interface {
	Create(assessment *entity.Assessment) error
	Select() ([]*entity.Assessment, error)
	Show(id *entity.InputID) (*entity.Assessment, error)
	Update(id *entity.InputID, assessment *entity.Assessment) error
	Delete(id *entity.InputID) error
}

type AssessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository(db *gorm.DB) IAssessmentRepository {
	return &AssessmentRepository{db: db}
}

func (a *AssessmentRepository) Create(assessment *entity.Assessment) error {
	res := a.db.Create(assessment)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (a *AssessmentRepository) Select() ([]*entity.Assessment, error) {
	var assessments []*entity.Assessment
	res := a.db.Find(&assessments)
	if res.Error != nil {
		return nil, res.Error
	}
	return assessments, nil
}

func (a *AssessmentRepository) Show(input *entity.InputID) (*entity.Assessment, error) {
	assessment := &entity.Assessment{}
	res := a.db.First(&assessment, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return assessment, nil
}

func (a *AssessmentRepository) Update(input *entity.InputID, assessment *entity.Assessment) error {
	res := a.db.Model(&assessment).Omit("ID").Where("id = ?", input.ID).Updates(entity.Assessment{Courses: assessment.Courses, Classrooms: assessment.Classrooms, StartDate: assessment.StartDate, EndDate: assessment.EndDate, Quiz: assessment.Quiz})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *AssessmentRepository) Delete(input *entity.InputID) error {
	assessment := &entity.Assessment{}

	res := a.db.First(&assessment, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}

	res = a.db.Delete(&assessment, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
