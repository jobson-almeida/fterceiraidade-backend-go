package integration_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"
	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/repository"
	"github.com/jobson-almeida/fterceiraidade-backend-go/tests/testhelpers"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AssessmentRepoTestSuite struct {
	suite.Suite
	container            *testhelpers.DatabaseContainer
	ctx                  context.Context
	conn                 *gorm.DB
	courseRepository     *repository.CourseRepository
	classroomRepository  *repository.ClassroomRepository
	questionRepository   *repository.QuestionRepository
	assessmentRepository *repository.AssessmentRepository
	course               string
	classroom            *entity.Classroom
	question             *entity.Question
	assessment           *entity.Assessment
	quiz                 []*entity.Quiz
	// created    bool
	// showed     bool
	// updated    bool
}

func (suite *AssessmentRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "assessments.sql")
	if err != nil {
		log.Fatal(err)
	}
	suite.container = container

	err = godotenv.Load("../../.env")
	if err != nil {
		panic(err.Error())
	}
	dialect := os.Getenv("DIALECT")
	suite.conn, err = gorm.Open(pg.New(pg.Config{DriverName: dialect, DSN: suite.container.ConnectionString}))
	if err != nil {
		panic(err.Error())
	}
}

func (suite *AssessmentRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *AssessmentRepoTestSuite) BeforeTest(suiteName, testName string) {
	t := suite.T()

	if testName == "TestCreateAssessment" {
		t.Run("create course", func(t *testing.T) {
			repository := repository.NewCourseRepository(suite.conn)
			suite.courseRepository = repository

			course, err := entity.NewCourse(
				"Name",
				"Description",
				"/image/image.png",
			)
			if err != nil {
				panic(err.Error())
			}
			suite.course = course.ID

			err = suite.courseRepository.Create(course)
			assert.NoError(t, err)
			assert.NotNil(t, course)
		})

		t.Run("create classroom", func(t *testing.T) {
			repository := repository.NewClassroomRepository(suite.conn)
			suite.classroomRepository = repository

			classroom, err := entity.NewClassroom(
				"Name",
				"Description",
				suite.course,
			)
			if err != nil {
				panic(err.Error())
			}

			err = suite.classroomRepository.Create(classroom)
			assert.NoError(t, err)
			assert.NotNil(t, classroom)
			suite.classroom = classroom
		})

		t.Run("create question", func(t *testing.T) {
			repository := repository.NewQuestionRepository(suite.conn)
			suite.questionRepository = repository

			var str_img = "/image/image.png"
			image := &str_img
			var str_answer = "Answer"
			answer := &str_answer
			alternatives := pq.StringArray{"(a) alternative a", "(b) alternative b", "(c) alternative c"}

			question, err := entity.NewQuestion(
				"Questioning",
				"Type",
				image,
				alternatives,
				answer,
				"Discipline",
			)
			if err != nil {
				panic(err.Error())
			}

			err = suite.questionRepository.Create(question)
			assert.NoError(t, err)
			assert.NotNil(t, question)
			suite.question = question
		})
	}
}

func (suite *AssessmentRepoTestSuite) AfterTest(_, _ string) {
}

func (suite *AssessmentRepoTestSuite) TestCreateAssessment() {
	t := suite.T()
	repository := repository.NewAssessmentRepository(suite.conn)
	suite.assessmentRepository = repository

	var quiz []*entity.Quiz
	quiz = append(quiz, &entity.Quiz{
		ID:    suite.question.ID,
		Value: 1,
	})
	courses := pq.StringArray{suite.course}
	classrooms := pq.StringArray{suite.classroom.ID}

	startDate, _ := time.Parse(time.DateOnly, "2024-01-01")
	endDate, _ := time.Parse(time.DateOnly, "2024-01-01")

	assessment, err := entity.NewAssessment(
		"Description",
		courses,
		classrooms,
		startDate,
		endDate,
		quiz,
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.assessmentRepository.Create(assessment)
	assert.NoError(t, err)
	assert.NotNil(t, assessment)

	suite.assessment = assessment
	suite.quiz = quiz
}

func (suite *AssessmentRepoTestSuite) TestShowAssessment() {
	t := suite.T()
	repository := repository.NewAssessmentRepository(suite.conn)
	suite.assessmentRepository = repository

	id, err := entity.NewInputID(suite.assessment.ID)
	assert.NoError(t, err)

	currentAssessment, err := suite.assessmentRepository.Show(id)
	assert.NoError(t, err)
	assert.NotNil(t, currentAssessment)
	assert.Equal(t, "Description", currentAssessment.Description)
	assert.ElementsMatch(t, suite.assessment.Courses, currentAssessment.Courses)
	assert.ElementsMatch(t, suite.assessment.Classrooms, currentAssessment.Classrooms)
	assert.Equal(t, suite.assessment.StartDate, currentAssessment.StartDate.In(time.UTC))
	assert.Equal(t, suite.assessment.EndDate, currentAssessment.EndDate.In(time.UTC))
	assert.ElementsMatch(t, suite.assessment.Quiz, currentAssessment.Quiz)
}

func TestAssessmentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AssessmentRepoTestSuite))
}
