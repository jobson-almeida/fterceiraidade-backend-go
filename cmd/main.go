package main

import (
	"net/http"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/database"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/routes"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/usecase"
	_ "github.com/lib/pq"

	handlers "github.com/jobson-almeida/fterceiraidade-backend-go/internal/app"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath
func main() {
	db := *database.Connection()

	repositoryCourse := repository.NewCourseRepository(&db)
	createCourse := usecase.NewCreateCourse(repositoryCourse)
	selectCourses := usecase.NewSelectCourses(repositoryCourse)
	showCourse := usecase.NewShowCourse(repositoryCourse)
	updateCourse := usecase.NewUpdateCourse(repositoryCourse)
	deleteCourse := usecase.NewDeleteCourse(repositoryCourse)
	courseHandlers := handlers.NewCourseHandlers(createCourse, selectCourses, showCourse, updateCourse, deleteCourse)

	repositoryTeacher := repository.NewTeacherRepository(&db)
	createTeacher := usecase.NewCreateTeacher(repositoryTeacher)
	selectTeachers := usecase.NewSelectTeachers(repositoryTeacher)
	showTeacher := usecase.NewShowTeacher(repositoryTeacher)
	updateTeacher := usecase.NewUpdateTeacher(repositoryTeacher)
	deleteTeacher := usecase.NewDeleteTeacher(repositoryTeacher)
	teacherHandlers := handlers.NewTeacherHandlers(createTeacher, selectTeachers, showTeacher, updateTeacher, deleteTeacher)

	repositoryStudent := repository.NewStudentRepository(&db)
	createStudent := usecase.NewCreateStudent(repositoryStudent)
	selectStudents := usecase.NewSelectStudents(repositoryStudent)
	showStudent := usecase.NewShowStudent(repositoryStudent)
	updateStudent := usecase.NewUpdateStudent(repositoryStudent)
	deleteStudent := usecase.NewDeleteStudent(repositoryStudent)
	studentHandlers := handlers.NewStudentHandlers(createStudent, selectStudents, showStudent, updateStudent, deleteStudent)

	repositoryQuestion := repository.NewQuestionRepository(&db)
	createQuestion := usecase.NewCreateQuestion(repositoryQuestion)
	selectQuestions := usecase.NewSelectQuestions(repositoryQuestion)
	showQuestion := usecase.NewShowQuestion(repositoryQuestion)
	updateQuestion := usecase.NewUpdateQuestion(repositoryQuestion)
	deleteQuestion := usecase.NewDeleteQuestion(repositoryQuestion)
	questionHandlers := handlers.NewQuestionHandlers(createQuestion, selectQuestions, showQuestion, updateQuestion, deleteQuestion)

	repositoryClassroom := repository.NewClassroomRepository(&db)
	createClassroom := usecase.NewCreateClassroom(repositoryClassroom)
	selectClassrooms := usecase.NewSelectClassrooms(repositoryClassroom)
	showClassroom := usecase.NewShowClassroom(repositoryClassroom)
	updateClassroom := usecase.NewUpdateClassroom(repositoryClassroom)
	deleteClassroom := usecase.NewDeleteClassroom(repositoryClassroom)
	classroomHandlers := handlers.NewClassroomHandlers(createClassroom, selectClassrooms, showClassroom, updateClassroom, deleteClassroom)

	repositoryAssessment := repository.NewAssessmentRepository(&db)
	createAssessment := usecase.NewCreateAssessment(repositoryAssessment)
	selectAssessments := usecase.NewSelectAssessments(repositoryAssessment)
	showAssessment := usecase.NewShowAssessment(repositoryAssessment)
	updateAssessment := usecase.NewUpdateAssessment(repositoryAssessment)
	deleteAssessment := usecase.NewDeleteAssessment(repositoryAssessment)
	assessmentHandlers := handlers.NewAssessmentHandlers(createAssessment, selectAssessments, showAssessment, updateAssessment, deleteAssessment)

	r := routes.Router(courseHandlers, studentHandlers, teacherHandlers, questionHandlers, classroomHandlers, assessmentHandlers)

	http.ListenAndServe(":8888", r)
}
