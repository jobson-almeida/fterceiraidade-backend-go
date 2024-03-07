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

type TeacherRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.TeacherRepository
	ctx        context.Context
	teacher    *entity.InputID
}

func (suite *TeacherRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "teachers.sql")
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
	repository := repository.NewTeacherRepository(conn)

	conn.AutoMigrate(
		entity.Teacher{},
	)
	suite.repository = repository
}

func (suite *TeacherRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *TeacherRepoTestSuite) TestCreateTeacher() {
	t := suite.T()

	address := entity.DetailsAddress{City: "City", State: "State", Street: "Street"}
	teacher, err := entity.NewTeacher(
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

	err = suite.repository.Create(teacher)
	assert.NoError(t, err)

	suite.teacher, err = entity.NewInputID(teacher.ID)
	assert.NoError(t, err)
}

func (suite *TeacherRepoTestSuite) TestShowTeacher() {
	t := suite.T()

	currentTeacher, err := suite.repository.Show(suite.teacher)
	assert.NoError(t, err)
	assert.NotNil(t, currentTeacher)
	assert.Equal(t, "/avatar/avatar.png", currentTeacher.Avatar)
	assert.Equal(t, "Firstname", currentTeacher.Firstname)
	assert.Equal(t, "Lastname", currentTeacher.Lastname)
	assert.Equal(t, "email@email.com", currentTeacher.Email)
	assert.Equal(t, "+998812345678", currentTeacher.Phone)
	assert.Equal(t, "City", currentTeacher.Address.City)
	assert.Equal(t, "State", currentTeacher.Address.State)
	assert.Equal(t, "Street", currentTeacher.Address.Street)
}

func TestTeacherRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TeacherRepoTestSuite))
}
