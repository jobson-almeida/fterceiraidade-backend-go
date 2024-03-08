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
	question   *entity.InputID
	created    bool
	//showed     bool
	//updated    bool
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

	suite.question, err = entity.NewInputID(question.ID)
	assert.NoError(t, err)
	suite.created = true
}

func TestQuestionRepoTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionRepoTestSuite))
}
