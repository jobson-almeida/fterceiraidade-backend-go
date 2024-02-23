package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"

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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.CreateStudent.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *StudentHandlers) SelectStudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectStudent.Execute()
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

func (c *StudentHandlers) ShowStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := c.ShowStudent.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *StudentHandlers) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowStudent.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var student dto.UpdateStudentInput
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.UpdateStudent.Execute(input, student)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *StudentHandlers) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowStudent.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.DeleteStudent.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
