package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	"github.com/jobson-almeida/fterceiraidade-backend-go/util"

	"github.com/go-chi/chi"
)

type QuestionHandlers struct {
	CreateQuestion *usecase.CreateQuestion
	SelectQuestion *usecase.SelectQuestions
	ShowQuestion   *usecase.ShowQuestion
	UpdateQuestion *usecase.UpdateQuestion
	DeleteQuestion *usecase.DeleteQuestion
}

type IQuestionHandlers interface {
	CreateQuestionHandler(w http.ResponseWriter, r *http.Request)
	SelectQuestionsHandler(w http.ResponseWriter, r *http.Request)
	ShowQuestionHandler(w http.ResponseWriter, r *http.Request)
	UpdateQuestionHandler(w http.ResponseWriter, r *http.Request)
	DeleteQuestionHandler(w http.ResponseWriter, r *http.Request)
}

func NewQuestionHandlers(createQuestion *usecase.CreateQuestion, selectQuestion *usecase.SelectQuestions, showQuestion *usecase.ShowQuestion,
	updateQuestion *usecase.UpdateQuestion, deleteQuestion *usecase.DeleteQuestion) IQuestionHandlers {
	return &QuestionHandlers{
		CreateQuestion: createQuestion,
		SelectQuestion: selectQuestion,
		ShowQuestion:   showQuestion,
		UpdateQuestion: updateQuestion,
		DeleteQuestion: deleteQuestion,
	}
}

func (c *QuestionHandlers) CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.QuestionInput
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.CreateQuestion.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *QuestionHandlers) SelectQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output, err := c.SelectQuestion.Execute()
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	if len(output) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *QuestionHandlers) ShowQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	output, err := c.ShowQuestion.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (c *QuestionHandlers) UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowQuestion.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	var question dto.UpdateQuestionInput
	err = json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.UpdateQuestion.Execute(input, question)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *QuestionHandlers) DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	_, err := c.ShowQuestion.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}

	err = c.DeleteQuestion.Execute(input)
	if err != nil {
		s, e := util.Error(err)
		w.WriteHeader(s)
		w.Write([]byte(e))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
