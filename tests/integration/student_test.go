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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.StudentRepository
	ctx        context.Context
	student    *entity.InputID
	created    bool
	//showed     bool
	//updated    bool
}

func (suite *StudentRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "students.sql")
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
	repository := repository.NewStudentRepository(conn)

	conn.AutoMigrate(
		entity.Student{},
	)
	suite.repository = repository
}

func (suite *StudentRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *StudentRepoTestSuite) AfterTest(_, _ string) {
}

func (suite *StudentRepoTestSuite) TestCreateStudent() {
	t := suite.T()

	address := entity.DetailsAddress{City: "City", State: "State", Street: "Street"}
	student, err := entity.NewStudent(
		"/avatar/avatar.png",
		"Firstname",
		"Lastname",
		"email@email.com",
		"+998812345678",
		address,
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.repository.Create(student)
	assert.NoError(t, err)

	suite.student, err = entity.NewInputID(student.ID)
	assert.NoError(t, err)
	suite.created = true
}

func TestStudentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepoTestSuite))
}
