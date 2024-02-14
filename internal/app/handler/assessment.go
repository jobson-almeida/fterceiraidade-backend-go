package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/lib/pq"

	"github.com/go-chi/chi"
)

type AssessmentHandlers struct {
	CreateAssessment *usecase.CreateAssessment
	SelectAssessment *usecase.SelectAssessments
	ShowAssessment   *usecase.ShowAssessment
	UpdateAssessment *usecase.UpdateAssessment
	DeleteAssessment *usecase.DeleteAssessment
	ShowCourse       *usecase.ShowCourse
	ShowClassroom    *usecase.ShowClassroom
}

type IAssessmentHandlers interface {
	CreateAssessmentHandler(w http.ResponseWriter, r *http.Request)
	SelectAssessmentsHandler(w http.ResponseWriter, r *http.Request)
	ShowAssessmentHandler(w http.ResponseWriter, r *http.Request)
	UpdateAssessmentHandler(w http.ResponseWriter, r *http.Request)
	DeleteAssessmentHandler(w http.ResponseWriter, r *http.Request)
	showCourseHandler(w http.ResponseWriter, r *http.Request, s pq.StringArray) error
	showClassroomHandler(w http.ResponseWriter, r *http.Request, s pq.StringArray) error
}

func (a *AssessmentHandlers) showClassroomHandler(w http.ResponseWriter, r *http.Request, c pq.StringArray) error {
	for _, r := range c {
		input := dto.IDInput{ID: r}
		_, err := a.ShowClassroom.Execute(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AssessmentHandlers) showCourseHandler(w http.ResponseWriter, r *http.Request, c pq.StringArray) error {
	for _, r := range c {
		input := dto.IDInput{ID: r}
		_, err := a.ShowCourse.Execute(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewAssessmentHandlers(createAssessment *usecase.CreateAssessment, selectAssessment *usecase.SelectAssessments, showAssessment *usecase.ShowAssessment,
	updateAssessment *usecase.UpdateAssessment, deleteAssessment *usecase.DeleteAssessment, showCourse *usecase.ShowCourse, showClassroom *usecase.ShowClassroom) IAssessmentHandlers {
	return &AssessmentHandlers{
		CreateAssessment: createAssessment,
		SelectAssessment: selectAssessment,
		ShowAssessment:   showAssessment,
		UpdateAssessment: updateAssessment,
		DeleteAssessment: deleteAssessment,
		ShowCourse:       showCourse,
		ShowClassroom:    showClassroom,
	}
}

func (a *AssessmentHandlers) CreateAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.AssessmentInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.CreateAssessment.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (a *AssessmentHandlers) SelectAssessmentsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := a.SelectAssessment.Execute()

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

func (a *AssessmentHandlers) ShowAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	output, err := a.ShowAssessment.Execute(input)
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

func (a *AssessmentHandlers) UpdateAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := a.ShowAssessment.Execute(input)
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

	var assessment dto.UpdateAssessmentInput
	err = json.NewDecoder(r.Body).Decode(&assessment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	/*
		err = a.showCourseHandler(w, r, assessment.Courses)
		if err != nil {
			if strings.TrimSpace(err.Error()) == "record not found" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("course not found"))
				return
			}
		}
	*/
	err = a.showClassroomHandler(w, r, assessment.Classrooms)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("classroom not found"))
			return
		} else {
			_, after, _ := strings.Cut(err.Error(), "pq: ")
			//e := strings.SplitN(string(err.Error()), "pq: ", 1)[0]
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(after))
			return
		}
	}

	err = a.UpdateAssessment.Execute(input, assessment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *AssessmentHandlers) DeleteAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := a.ShowAssessment.Execute(input)
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

	err = a.DeleteAssessment.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
