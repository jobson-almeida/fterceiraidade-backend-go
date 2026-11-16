package repository

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"gorm.io/gorm"
)

type IChartDataRepository interface {
	Select() ([]*entity.ChartData, error)
}

type ChartDataRepository struct {
	db *gorm.DB
}

func NewChartDataRepository(db *gorm.DB) *ChartDataRepository {
	return &ChartDataRepository{db: db}
}

func (s *ChartDataRepository) Select() ([]*entity.ChartData, error) {
	var ChartDatas []*entity.ChartData
	res := s.db.Find(&ChartDatas)
	if res.Error != nil {
		return nil, res.Error
	}
	return ChartDatas, nil
}
