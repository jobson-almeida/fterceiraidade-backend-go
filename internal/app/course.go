package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	usecaseCourse "github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

	"github.com/go-chi/chi"
)

// adapter
type CourseHandlers struct {
	CreateCourse *usecaseCourse.CreateCourse
	SelectCourse *usecaseCourse.SelectCourses
	ShowCourse   *usecaseCourse.ShowCourse
	UpdateCourse *usecaseCourse.UpdateCourse
	DeleteCourse *usecaseCourse.DeleteCourse
}

type ICourseHandlers interface {
	CreateCourseHandlers(w http.ResponseWriter, r *http.Request)
	SelectCoursesHandlers(w http.ResponseWriter, r *http.Request)
	ShowCourseHandlers(w http.ResponseWriter, r *http.Request)
	UpdateCourseHandlers(w http.ResponseWriter, r *http.Request)
	DeleteCourseHandlers(w http.ResponseWriter, r *http.Request)
}

func NewCourseHandlers(createCourse *usecaseCourse.CreateCourse, selectCourse *usecaseCourse.SelectCourses, showCourse *usecaseCourse.ShowCourse,
	updateCourse *usecaseCourse.UpdateCourse, deleteCourse *usecaseCourse.DeleteCourse) ICourseHandlers {
	return &CourseHandlers{
		CreateCourse: createCourse,
		SelectCourse: selectCourse,
		ShowCourse:   showCourse,
		UpdateCourse: updateCourse,
		DeleteCourse: deleteCourse,
	}
}

func (c *CourseHandlers) CreateCourseHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *CourseHandlers) SelectCoursesHandlers(w http.ResponseWriter, r *http.Request) {
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

func (c *CourseHandlers) ShowCourseHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowCourse.Execute(input)
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

func (c *CourseHandlers) UpdateCourseHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowCourse.Execute(input)
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

func (c *CourseHandlers) DeleteCourseHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowCourse.Execute(input)
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

	err = c.DeleteCourse.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
