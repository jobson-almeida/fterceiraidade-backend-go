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
type CourseHandlers struct {
	CreateCourse *usecase.CreateCourse
	SelectCourse *usecase.SelectCourses
	ShowCourse   *usecase.ShowCourse
	UpdateCourse *usecase.UpdateCourse
	DeleteCourse *usecase.DeleteCourse
}

type ICourseHandlers interface {
	CreateCourseHandler(w http.ResponseWriter, r *http.Request)
	SelectCoursesHandler(w http.ResponseWriter, r *http.Request)
	ShowCourseHandler(w http.ResponseWriter, r *http.Request)
	UpdateCourseHandler(w http.ResponseWriter, r *http.Request)
	DeleteCourseHandler(w http.ResponseWriter, r *http.Request)
}

func NewCourseHandlers(createCourse *usecase.CreateCourse, selectCourse *usecase.SelectCourses, showCourse *usecase.ShowCourse,
	updateCourse *usecase.UpdateCourse, deleteCourse *usecase.DeleteCourse) ICourseHandlers {
	return &CourseHandlers{
		CreateCourse: createCourse,
		SelectCourse: selectCourse,
		ShowCourse:   showCourse,
		UpdateCourse: updateCourse,
		DeleteCourse: deleteCourse,
	}
}

func (c *CourseHandlers) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CourseInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.CreateCourse.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *CourseHandlers) SelectCoursesHandler(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectCourse.Execute()

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

func (c *CourseHandlers) ShowCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	output, err := c.ShowCourse.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("course not found"))
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

func (c *CourseHandlers) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowCourse.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("course not found"))
			return
		} else {
			_, after, _ := strings.Cut(err.Error(), "pq: ")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(after))
			return
		}
	}

	var course dto.UpdateCourseInput
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.UpdateCourse.Execute(input, course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *CourseHandlers) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowCourse.Execute(input)
	if err != nil {
		if strings.TrimSpace(err.Error()) == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("course not found"))
			return
		} else {
			_, after, _ := strings.Cut(err.Error(), "pq: ")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(after))
			return
		}
	}

	err = c.DeleteCourse.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
