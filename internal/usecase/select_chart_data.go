package usecase

import (
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
)

type SelectChartData struct {
	ChartDataRepository repository.IChartDataRepository
}

func NewSelectChartData(repository repository.IChartDataRepository) *SelectChartData {
	return &SelectChartData{ChartDataRepository: repository}
}

func (s *SelectChartData) Execute() ([]*dto.ChartDataOutput, error) {
	res, err := s.ChartDataRepository.Select()
	if err != nil {
		return []*dto.ChartDataOutput{}, err
	}

	var output []*dto.ChartDataOutput
	for _, r := range res {
		output = append(output, &dto.ChartDataOutput{
			Name:  r.Name,
			Value: r.Value,
		})
	}
	return output, nil
}
