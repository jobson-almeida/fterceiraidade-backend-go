package repository

import (
	"errors"
	"fterceiraidade-backend-go/internal/entity"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	Create(question *entity.Question) error
	Select() ([]*entity.Question, error)
	Show(id *entity.InputID) (*entity.Question, error)
	Update(id *entity.InputID, question *entity.Question) error
	Delete(id *entity.InputID) error
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{db: db}
}

func (q *QuestionRepository) Create(question *entity.Question) error {
	res := q.db.Create(question)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no insert effected")
	}
	return nil
}

func (q *QuestionRepository) Select() ([]*entity.Question, error) {
	var questions []*entity.Question
	res := q.db.Find(&questions)
	if res.Error != nil {
		return nil, res.Error
	}
	return questions, nil
}

func (q *QuestionRepository) Show(input *entity.InputID) (*entity.Question, error) {
	question := &entity.Question{}
	res := q.db.First(&question, "id = ?", input.ID)
	if res.Error != nil {
		return nil, res.Error
	}
	return question, nil
}

func (q *QuestionRepository) Update(input *entity.InputID, question *entity.Question) error {
	res := q.db.Model(&question).Omit("ID").Where("id = ?", input.ID).Updates(entity.Question{
		Questioning: question.Questioning, Type: question.Type, Image: question.Image,
		Alternatives: question.Alternatives, Answer: question.Answer, Discipline: question.Discipline,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (q *QuestionRepository) Delete(input *entity.InputID) error {
	question := &entity.Question{}
	res := q.db.Delete(&question, "id = ?", input.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
