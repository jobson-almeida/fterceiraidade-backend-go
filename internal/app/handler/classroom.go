package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

// adapter
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
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.CreateClassroom.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *ClassroomHandlers) SelectClassroomsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectClassroom.Execute()

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

func (c *ClassroomHandlers) ShowClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowClassroom.Execute(input)
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

func (c *ClassroomHandlers) UpdateClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowClassroom.Execute(input)
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

	var classroom dto.UpdateClassroomInput
	err = json.NewDecoder(r.Body).Decode(&classroom)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.UpdateClassroom.Execute(input, classroom)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *ClassroomHandlers) DeleteClassroomHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowClassroom.Execute(input)
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

	err = c.DeleteClassroom.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
