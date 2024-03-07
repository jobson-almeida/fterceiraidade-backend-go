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

type ClassroomRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.ClassroomRepository
	ctx        context.Context
	course     string
	classroom  *entity.InputID
	// created    bool
	// showed     bool
	// updated    bool
}

func (suite *ClassroomRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "classrooms.sql")
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
	repository := repository.NewClassroomRepository(conn)

	conn.AutoMigrate(
		entity.Classroom{},
	)
	suite.repository = repository
}

func (suite *ClassroomRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *ClassroomRepoTestSuite) AfterTest(_, _ string) {
}

func (suite *ClassroomRepoTestSuite) TestCreateClassroom() {
	t := suite.T()

	course, err := entity.NewCourse(
		"Name",
		"Description",
		"/image/image.png",
	)
	if err != nil {
		panic(err.Error())
	}

	classroom, err := entity.NewClassroom(
		"Name",
		"Description",
		course.ID,
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.repository.Create(classroom)
	assert.NoError(t, err)
	assert.NotNil(t, classroom)

	suite.classroom, err = entity.NewInputID(classroom.ID)
	assert.NoError(t, err)

	suite.course = course.ID
}

func (suite *ClassroomRepoTestSuite) TestShowClassroom() {
	t := suite.T()

	currentClassroom, err := suite.repository.Show(suite.classroom)
	assert.NoError(t, err)
	assert.NotNil(t, currentClassroom)
	assert.Equal(t, "Name", currentClassroom.Name)
	assert.Equal(t, "Description", currentClassroom.Description)
	assert.Equal(t, suite.course, currentClassroom.Course)
}

func TestClassroomRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ClassroomRepoTestSuite))
}
