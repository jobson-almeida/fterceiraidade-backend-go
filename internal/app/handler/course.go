package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"

	"github.com/go-chi/chi"
)

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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.CreateCourse.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CourseHandlers) SelectCoursesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectCourse.Execute()
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

func (c *CourseHandlers) ShowCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := c.ShowCourse.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *CourseHandlers) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowCourse.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var course dto.UpdateCourseInput
	err = json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.UpdateCourse.Execute(input, course)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *CourseHandlers) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowCourse.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.DeleteCourse.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
