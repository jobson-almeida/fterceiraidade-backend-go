package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	usecaseTeacher "github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

type TeacherHandlers struct {
	CreateTeacher *usecaseTeacher.CreateTeacher
	SelectTeacher *usecaseTeacher.SelectTeachers
	ShowTeacher   *usecaseTeacher.ShowTeacher
	UpdateTeacher *usecaseTeacher.UpdateTeacher
	DeleteTeacher *usecaseTeacher.DeleteTeacher
}

type ITeacherHandlers interface {
	CreateTeacherHandlers(w http.ResponseWriter, r *http.Request)
	SelectTeachersHandlers(w http.ResponseWriter, r *http.Request)
	ShowTeacherHandlers(w http.ResponseWriter, r *http.Request)
	UpdateTeacherHandlers(w http.ResponseWriter, r *http.Request)
	DeleteTeacherHandlers(w http.ResponseWriter, r *http.Request)
}

func NewTeacherHandlers(createTeacher *usecaseTeacher.CreateTeacher, selectTeacher *usecaseTeacher.SelectTeachers, showTeacher *usecaseTeacher.ShowTeacher,
	updateTeacher *usecaseTeacher.UpdateTeacher, deleteTeacher *usecaseTeacher.DeleteTeacher) ITeacherHandlers {
	return &TeacherHandlers{
		CreateTeacher: createTeacher,
		SelectTeacher: selectTeacher,
		ShowTeacher:   showTeacher,
		UpdateTeacher: updateTeacher,
		DeleteTeacher: deleteTeacher,
	}
}

func (c *TeacherHandlers) CreateTeacherHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.TeacherInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.CreateTeacher.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *TeacherHandlers) SelectTeachersHandlers(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectTeacher.Execute()

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

func (c *TeacherHandlers) ShowTeacherHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowTeacher.Execute(input)
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

func (c *TeacherHandlers) UpdateTeacherHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowTeacher.Execute(input)
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

	var teacher dto.UpdateTeacherInput
	err = json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.UpdateTeacher.Execute(input, teacher)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *TeacherHandlers) DeleteTeacherHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowTeacher.Execute(input)
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

	err = c.DeleteTeacher.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
