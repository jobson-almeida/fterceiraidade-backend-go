package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	usecaseClassroom "github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

// adapter
type ClassroomHandlers struct {
	CreateClassroom *usecaseClassroom.CreateClassroom
	SelectClassroom *usecaseClassroom.SelectClassrooms
	ShowClassroom   *usecaseClassroom.ShowClassroom
	UpdateClassroom *usecaseClassroom.UpdateClassroom
	DeleteClassroom *usecaseClassroom.DeleteClassroom
}

type IClassroomHandlers interface {
	CreateClassroomHandlers(w http.ResponseWriter, r *http.Request)
	SelectClassroomsHandlers(w http.ResponseWriter, r *http.Request)
	ShowClassroomHandlers(w http.ResponseWriter, r *http.Request)
	UpdateClassroomHandlers(w http.ResponseWriter, r *http.Request)
	DeleteClassroomHandlers(w http.ResponseWriter, r *http.Request)
}

func NewClassroomHandlers(createClassroom *usecaseClassroom.CreateClassroom, selectClassroom *usecaseClassroom.SelectClassrooms, showClassroom *usecaseClassroom.ShowClassroom,
	updateClassroom *usecaseClassroom.UpdateClassroom, deleteClassroom *usecaseClassroom.DeleteClassroom) IClassroomHandlers {
	return &ClassroomHandlers{
		CreateClassroom: createClassroom,
		SelectClassroom: selectClassroom,
		ShowClassroom:   showClassroom,
		UpdateClassroom: updateClassroom,
		DeleteClassroom: deleteClassroom,
	}
}

func (c *ClassroomHandlers) CreateClassroomHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *ClassroomHandlers) SelectClassroomsHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *ClassroomHandlers) ShowClassroomHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *ClassroomHandlers) UpdateClassroomHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *ClassroomHandlers) DeleteClassroomHandlers(w http.ResponseWriter, r *http.Request) {
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
