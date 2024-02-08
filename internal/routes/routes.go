package routes

import (
	"fterceiraidade-backend-go/internal/app"
	"net/http"

	//"github.com/jobson-almeida/fterceiraidade-backend-go/internal/app"

	"github.com/go-chi/chi"
)

func Router(
	courseHandlers app.ICourseHandlers,
	studentHandlers app.IStudentHandlers,
	teacherHandlers app.ITeacherHandlers,
	questionHandlers app.IQuestionHandlers,
	classroomHandlers app.IClassroomHandlers,
	assessmentHandlers app.IAssessmentHandlers,
) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi :)"))
	})

	r.Route("/courses", func(r chi.Router) {
		r.Get("/", courseHandlers.SelectCoursesHandlers)
		r.Post("/", courseHandlers.CreateCourseHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", courseHandlers.ShowCourseHandlers)
			r.Put("/", courseHandlers.UpdateCourseHandlers)
			r.Delete("/", courseHandlers.DeleteCourseHandlers)
		})
	})

	r.Route("/students", func(r chi.Router) {
		r.Get("/", studentHandlers.SelectStudentsHandlers)
		r.Post("/", studentHandlers.CreateStudentHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", studentHandlers.ShowStudentHandlers)
			r.Put("/", studentHandlers.UpdateStudentHandlers)
			r.Delete("/", studentHandlers.DeleteStudentHandlers)
		})
	})

	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", teacherHandlers.SelectTeachersHandlers)
		r.Post("/", teacherHandlers.CreateTeacherHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", teacherHandlers.ShowTeacherHandlers)
			r.Put("/", teacherHandlers.UpdateTeacherHandlers)
			r.Delete("/", teacherHandlers.DeleteTeacherHandlers)
		})
	})

	r.Route("/questions", func(r chi.Router) {
		r.Get("/", questionHandlers.SelectQuestionsHandlers)
		r.Post("/", questionHandlers.CreateQuestionHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", questionHandlers.ShowQuestionHandlers)
			r.Put("/", questionHandlers.UpdateQuestionHandlers)
			r.Delete("/", questionHandlers.DeleteQuestionHandlers)
		})
	})

	r.Route("/classrooms", func(r chi.Router) {
		r.Get("/", classroomHandlers.SelectClassroomsHandlers)
		r.Post("/", classroomHandlers.CreateClassroomHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", classroomHandlers.ShowClassroomHandlers)
			r.Put("/", classroomHandlers.UpdateClassroomHandlers)
			r.Delete("/", classroomHandlers.DeleteClassroomHandlers)
		})
	})

	r.Route("/assessments", func(r chi.Router) {
		r.Get("/", assessmentHandlers.SelectAssessmentsHandlers)
		r.Post("/", assessmentHandlers.CreateAssessmentHandlers)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", assessmentHandlers.ShowAssessmentHandlers)
			r.Put("/", assessmentHandlers.UpdateAssessmentHandlers)
			r.Delete("/", assessmentHandlers.DeleteAssessmentHandlers)
		})
	})

	return r
}
