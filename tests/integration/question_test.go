package integration_test

import (
	"context"
	"log"
	"os"
	"testing"

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

type QuestionRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.QuestionRepository
	ctx        context.Context
	question   *entity.Question
	created    bool
	showed     bool
	updated    bool
}

func (suite *QuestionRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "questions.sql")
	if err != nil {
		log.Fatal(err)
	}
	suite.container = container

	err = godotenv.Load("../../.env")
	if err != nil {
		panic(err.Error())
	}
	dialect := os.Getenv("DIALECT")
	conn, err := gorm.Open(pg.New(pg.Config{DriverName: dialect, DSN: suite.container.ConnectionString}))
	if err != nil {
		panic(err.Error())
	}
	repository := repository.NewQuestionRepository(conn)

	conn.AutoMigrate(
		entity.Question{},
	)
	suite.repository = repository
}

func (suite *QuestionRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *QuestionRepoTestSuite) AfterTest(_, _ string) {
	if suite.created == true && suite.showed == true && suite.updated == true {
		t := suite.T()
		t.Run("delete question", func(t *testing.T) {
			id, err := entity.NewInputID(suite.question.ID)
			assert.NoError(t, err)

			err = suite.repository.Delete(id)
			assert.NoError(t, err)

			currentQuestion, err := suite.repository.Show(id)
			assert.Error(t, err)
			assert.Nil(t, currentQuestion)
		})
	}
}

func (suite *QuestionRepoTestSuite) TestCreateQuestion() {
	t := suite.T()

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

	err = suite.repository.Create(question)
	assert.NoError(t, err)
	assert.NotNil(t, question)

	suite.question = question
	suite.created = true
}

func (suite *QuestionRepoTestSuite) TestShowQuestion() {
	t := suite.T()

	id, err := entity.NewInputID(suite.question.ID)
	assert.NoError(t, err)

	currentQuestion, err := suite.repository.Show(id)
	assert.NoError(t, err)
	assert.NotNil(t, currentQuestion)
	assert.Equal(t, suite.question.Questioning, currentQuestion.Questioning)
	assert.Equal(t, suite.question.Type, currentQuestion.Type)
	assert.Equal(t, suite.question.Image, currentQuestion.Image)
	assert.ElementsMatch(t, suite.question.Alternatives, currentQuestion.Alternatives)
	assert.Equal(t, suite.question.Answer, currentQuestion.Answer)
	assert.Equal(t, suite.question.Discipline, currentQuestion.Discipline)
	suite.showed = true
}

func (suite *QuestionRepoTestSuite) TestUpdateQuestion() {
	t := suite.T()

	var str_img = "/image/new_image.png"
	image := &str_img
	var str_answer = "New Answer"
	answer := &str_answer
	alternatives := pq.StringArray{"(d) alternative d", "(e) alternative e", "(f) alternative f"}

	newQuestion, err := entity.NewQuestion(
		"New Questioning",
		"New Type",
		image,
		alternatives,
		answer,
		"New Discipline",
	)
	if err != nil {
		panic(err.Error())
	}

	id, err := entity.NewInputID(suite.question.ID)
	assert.NoError(t, err)

	err = suite.repository.Update(id, newQuestion)
	assert.NoError(t, err)
	assert.NotNil(t, newQuestion)
	assert.NotEqual(t, suite.question.Questioning, newQuestion.Questioning)
	assert.NotEqual(t, suite.question.Type, newQuestion.Type)
	assert.NotEqual(t, suite.question.Image, newQuestion.Image)
	assert.NotEqual(t, suite.question.Alternatives, newQuestion.Alternatives)
	assert.NotEqual(t, suite.question.Answer, newQuestion.Answer)
	assert.NotEqual(t, suite.question.Discipline, newQuestion.Discipline)
	suite.updated = true
}

func TestQuestionRepoTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionRepoTestSuite))
}
