package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/app/handler"
)

func Router(
	courseHandlers handler.ICourseHandlers,
	studentHandlers handler.IStudentHandlers,
	teacherHandlers handler.ITeacherHandlers,
	questionHandlers handler.IQuestionHandlers,
	classroomHandlers handler.IClassroomHandlers,
	assessmentHandlers handler.IAssessmentHandlers,
) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi :)"))
	})

	r.Route("/courses", func(r chi.Router) {
		r.Get("/", courseHandlers.SelectCoursesHandler)
		r.Post("/", courseHandlers.CreateCourseHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", courseHandlers.ShowCourseHandler)
			r.Put("/", courseHandlers.UpdateCourseHandler)
			r.Delete("/", courseHandlers.DeleteCourseHandler)
		})
	})

	r.Route("/students", func(r chi.Router) {
		r.Get("/", studentHandlers.SelectStudentsHandler)
		r.Post("/", studentHandlers.CreateStudentHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", studentHandlers.ShowStudentHandler)
			r.Put("/", studentHandlers.UpdateStudentHandler)
			r.Delete("/", studentHandlers.DeleteStudentHandler)
		})
	})

	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", teacherHandlers.SelectTeachersHandler)
		r.Post("/", teacherHandlers.CreateTeacherHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", teacherHandlers.ShowTeacherHandler)
			r.Put("/", teacherHandlers.UpdateTeacherHandler)
			r.Delete("/", teacherHandlers.DeleteTeacherHandler)
		})
	})

	r.Route("/questions", func(r chi.Router) {
		r.Get("/", questionHandlers.SelectQuestionsHandler)
		r.Post("/", questionHandlers.CreateQuestionHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", questionHandlers.ShowQuestionHandler)
			r.Put("/", questionHandlers.UpdateQuestionHandler)
			r.Delete("/", questionHandlers.DeleteQuestionHandler)
		})
	})

	r.Route("/classrooms", func(r chi.Router) {
		r.Get("/", classroomHandlers.SelectClassroomsHandler)
		r.Post("/", classroomHandlers.CreateClassroomHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", classroomHandlers.ShowClassroomHandler)
			r.Put("/", classroomHandlers.UpdateClassroomHandler)
			r.Delete("/", classroomHandlers.DeleteClassroomHandler)
		})
	})

	r.Route("/assessments", func(r chi.Router) {
		r.Get("/", assessmentHandlers.SelectAssessmentsHandler)
		r.Post("/", assessmentHandlers.CreateAssessmentHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", assessmentHandlers.ShowAssessmentHandler)
			r.Put("/", assessmentHandlers.UpdateAssessmentHandler)
			r.Delete("/", assessmentHandlers.DeleteAssessmentHandler)
		})
	})

	return r
}
