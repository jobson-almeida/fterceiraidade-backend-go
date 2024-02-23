package handler

import (
	"encoding/json"
	"net/http"

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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.CreateTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *TeacherHandlers) SelectTeachersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectTeacher.Execute()
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

func (c *TeacherHandlers) ShowTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := c.ShowTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *TeacherHandlers) UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var teacher dto.UpdateTeacherInput
	err = json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.UpdateTeacher.Execute(input, teacher)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *TeacherHandlers) DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
	}

	err = c.DeleteTeacher.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
	}
	w.WriteHeader(http.StatusNoContent)
}
