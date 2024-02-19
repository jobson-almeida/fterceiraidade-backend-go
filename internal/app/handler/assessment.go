package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"
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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = a.CreateAssessment.Execute(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AssessmentHandlers) SelectAssessmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := a.SelectAssessment.Execute()
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	if len(output) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (a *AssessmentHandlers) ShowAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := a.ShowAssessment.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (a *AssessmentHandlers) UpdateAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := a.ShowAssessment.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var assessment dto.UpdateAssessmentInput
	err = json.NewDecoder(r.Body).Decode(&assessment)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = a.showCourseHandler(w, r, assessment.Courses)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = a.showClassroomHandler(w, r, assessment.Classrooms)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = a.UpdateAssessment.Execute(input, assessment)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *AssessmentHandlers) DeleteAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := a.ShowAssessment.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = a.DeleteAssessment.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
