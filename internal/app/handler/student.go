package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

type StudentHandlers struct {
	CreateStudent *usecase.CreateStudent
	SelectStudent *usecase.SelectStudents
	ShowStudent   *usecase.ShowStudent
	UpdateStudent *usecase.UpdateStudent
	DeleteStudent *usecase.DeleteStudent
}

type IStudentHandlers interface {
	CreateStudentHandler(w http.ResponseWriter, r *http.Request)
	SelectStudentsHandler(w http.ResponseWriter, r *http.Request)
	ShowStudentHandler(w http.ResponseWriter, r *http.Request)
	UpdateStudentHandler(w http.ResponseWriter, r *http.Request)
	DeleteStudentHandler(w http.ResponseWriter, r *http.Request)
}

func NewStudentHandlers(createStudent *usecase.CreateStudent, selectStudent *usecase.SelectStudents, showStudent *usecase.ShowStudent,
	updateStudent *usecase.UpdateStudent, deleteStudent *usecase.DeleteStudent) IStudentHandlers {
	return &StudentHandlers{
		CreateStudent: createStudent,
		SelectStudent: selectStudent,
		ShowStudent:   showStudent,
		UpdateStudent: updateStudent,
		DeleteStudent: deleteStudent,
	}
}

func (c *StudentHandlers) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.StudentInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.CreateStudent.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *StudentHandlers) SelectStudentsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectStudent.Execute()

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

func (c *StudentHandlers) ShowStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowStudent.Execute(input)
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

func (c *StudentHandlers) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowStudent.Execute(input)
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

	var student dto.UpdateStudentInput
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = c.UpdateStudent.Execute(input, student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *StudentHandlers) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowStudent.Execute(input)
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

	err = c.DeleteStudent.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
