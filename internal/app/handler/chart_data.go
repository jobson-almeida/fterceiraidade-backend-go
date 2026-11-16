package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
)

type ChartDataHandlers struct {
	SelectChartData *usecase.SelectChartData
}

type IChartDataHandlers interface {
	SelectChartDatasHandler(w http.ResponseWriter, r *http.Request)
}

func NewChartDataHandlers(selectChartData *usecase.SelectChartData) IChartDataHandlers {
	return &ChartDataHandlers{
		SelectChartData: selectChartData,
	}
}

func (c *ChartDataHandlers) SelectChartDatasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectChartData.Execute()
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	if len(output) == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
