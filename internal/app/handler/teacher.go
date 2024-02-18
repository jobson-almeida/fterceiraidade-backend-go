package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"

	"github.com/go-chi/chi"
)

type TeacherHandlers struct {
	CreateTeacher *usecase.CreateTeacher
	SelectTeacher *usecase.SelectTeachers
	ShowTeacher   *usecase.ShowTeacher
	UpdateTeacher *usecase.UpdateTeacher
	DeleteTeacher *usecase.DeleteTeacher
}

type ITeacherHandlers interface {
	CreateTeacherHandler(w http.ResponseWriter, r *http.Request)
	SelectTeachersHandler(w http.ResponseWriter, r *http.Request)
	ShowTeacherHandler(w http.ResponseWriter, r *http.Request)
	UpdateTeacherHandler(w http.ResponseWriter, r *http.Request)
	DeleteTeacherHandler(w http.ResponseWriter, r *http.Request)
}

func NewTeacherHandlers(createTeacher *usecase.CreateTeacher, selectTeacher *usecase.SelectTeachers, showTeacher *usecase.ShowTeacher,
	updateTeacher *usecase.UpdateTeacher, deleteTeacher *usecase.DeleteTeacher) ITeacherHandlers {
	return &TeacherHandlers{
		CreateTeacher: createTeacher,
		SelectTeacher: selectTeacher,
		ShowTeacher:   showTeacher,
		UpdateTeacher: updateTeacher,
		DeleteTeacher: deleteTeacher,
	}
}

func (c *TeacherHandlers) CreateTeacherHandler(w http.ResponseWriter, r *http.Request) {
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

func (c *TeacherHandlers) SelectTeachersHandler(w http.ResponseWriter, r *http.Request) {
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

func (c *TeacherHandlers) ShowTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowTeacher.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("teacher not found"))
			return
		} else {
			_, after, _ := strings.Cut(err.Error(), "pq: ")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(after))
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *TeacherHandlers) UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		http.Error(w, e, s)
		return
	}

	var teacher dto.UpdateTeacherInput
	err = json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		s, e := util.Error(err)
		http.Error(w, e, s)
		return
	}

	err = c.UpdateTeacher.Execute(input, teacher)
	if err != nil {
		s, e := util.Error(err)
		http.Error(w, e, s)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *TeacherHandlers) DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowTeacher.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("teacher not found"))
			return
		} else {
			_, after, _ := strings.Cut(err.Error(), "pq: ")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(after))
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
