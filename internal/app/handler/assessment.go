package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

type AssessmentHandlers struct {
	CreateAssessment *usecase.CreateAssessment
	SelectAssessment *usecase.SelectAssessments
	ShowAssessment   *usecase.ShowAssessment
	UpdateAssessment *usecase.UpdateAssessment
	DeleteAssessment *usecase.DeleteAssessment
}

type IAssessmentHandlers interface {
	CreateAssessmentHandler(w http.ResponseWriter, r *http.Request)
	SelectAssessmentsHandler(w http.ResponseWriter, r *http.Request)
	ShowAssessmentHandler(w http.ResponseWriter, r *http.Request)
	UpdateAssessmentHandler(w http.ResponseWriter, r *http.Request)
	DeleteAssessmentHandler(w http.ResponseWriter, r *http.Request)
}

func NewAssessmentHandlers(createAssessment *usecase.CreateAssessment, selectAssessment *usecase.SelectAssessments, showAssessment *usecase.ShowAssessment,
	updateAssessment *usecase.UpdateAssessment, deleteAssessment *usecase.DeleteAssessment) IAssessmentHandlers {
	return &AssessmentHandlers{
		CreateAssessment: createAssessment,
		SelectAssessment: selectAssessment,
		ShowAssessment:   showAssessment,
		UpdateAssessment: updateAssessment,
		DeleteAssessment: deleteAssessment,
	}
}

func (c *AssessmentHandlers) CreateAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.AssessmentInput

	/*	err := decodeJSONBody(w, r, &input)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}*/

	//-------
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.CreateAssessment.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *AssessmentHandlers) SelectAssessmentsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectAssessment.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(output) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *AssessmentHandlers) ShowAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	output, err := c.ShowAssessment.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			json.Marshal([]string{})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *AssessmentHandlers) UpdateAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	/*
		_, err := c.ShowAssessment.Execute(input)
		if err != nil {
			if strings.TrimSpace(err.Error()) == "record not found" {
				w.WriteHeader(http.StatusNotFound)
				json.Marshal([]string{})
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		err = decodeJSONBody(w, r, &input)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

	*/
	_, err := c.ShowAssessment.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			json.Marshal([]string{})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	/*	var assessment dto.UpdateAssessmentInput
		err = json.NewDecoder(r.Body).Decode(&assessment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
	var assessment dto.UpdateAssessmentInput
	err = c.UpdateAssessment.Execute(input, assessment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *AssessmentHandlers) DeleteAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowAssessment.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			json.Marshal([]string{})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = c.DeleteAssessment.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
