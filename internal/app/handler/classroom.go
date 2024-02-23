package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"

	"github.com/go-chi/chi"
)

type ClassroomHandlers struct {
	CreateClassroom *usecase.CreateClassroom
	SelectClassroom *usecase.SelectClassrooms
	ShowClassroom   *usecase.ShowClassroom
	UpdateClassroom *usecase.UpdateClassroom
	DeleteClassroom *usecase.DeleteClassroom
}

type IClassroomHandlers interface {
	CreateClassroomHandler(w http.ResponseWriter, r *http.Request)
	SelectClassroomsHandler(w http.ResponseWriter, r *http.Request)
	ShowClassroomHandler(w http.ResponseWriter, r *http.Request)
	UpdateClassroomHandler(w http.ResponseWriter, r *http.Request)
	DeleteClassroomHandler(w http.ResponseWriter, r *http.Request)
}

func NewClassroomHandlers(createClassroom *usecase.CreateClassroom, selectClassroom *usecase.SelectClassrooms, showClassroom *usecase.ShowClassroom,
	updateClassroom *usecase.UpdateClassroom, deleteClassroom *usecase.DeleteClassroom) IClassroomHandlers {
	return &ClassroomHandlers{
		CreateClassroom: createClassroom,
		SelectClassroom: selectClassroom,
		ShowClassroom:   showClassroom,
		UpdateClassroom: updateClassroom,
		DeleteClassroom: deleteClassroom,
	}
}

func (c *ClassroomHandlers) CreateClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.ClassroomInput
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.CreateClassroom.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *ClassroomHandlers) SelectClassroomsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectClassroom.Execute()
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

func (c *ClassroomHandlers) ShowClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := c.ShowClassroom.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *ClassroomHandlers) UpdateClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowClassroom.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var classroom dto.UpdateClassroomInput
	err = json.NewDecoder(r.Body).Decode(&classroom)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.UpdateClassroom.Execute(input, classroom)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *ClassroomHandlers) DeleteClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowClassroom.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.DeleteClassroom.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
