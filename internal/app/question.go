package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/dto"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"

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
	CreateQuestionHandlers(w http.ResponseWriter, r *http.Request)
	SelectQuestionsHandlers(w http.ResponseWriter, r *http.Request)
	ShowQuestionHandlers(w http.ResponseWriter, r *http.Request)
	UpdateQuestionHandlers(w http.ResponseWriter, r *http.Request)
	DeleteQuestionHandlers(w http.ResponseWriter, r *http.Request)
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

func (c *QuestionHandlers) CreateQuestionHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.QuestionInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.CreateQuestion.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *QuestionHandlers) SelectQuestionsHandlers(w http.ResponseWriter, r *http.Request) {
	output, err := c.SelectQuestion.Execute()

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

func (c *QuestionHandlers) ShowQuestionHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")
	output, err := c.ShowQuestion.Execute(input)
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

func (c *QuestionHandlers) UpdateQuestionHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowQuestion.Execute(input)
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

	var question dto.UpdateQuestionInput
	err = json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.UpdateQuestion.Execute(input, question)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *QuestionHandlers) DeleteQuestionHandlers(w http.ResponseWriter, r *http.Request) {
	var input dto.IDInput
	input.ID = chi.URLParam(r, "id")

	_, err := c.ShowQuestion.Execute(input)
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

	err = c.DeleteQuestion.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
