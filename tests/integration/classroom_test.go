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
	created    bool
	showed     bool
	updated    bool
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
	if suite.created == true && suite.showed == true && suite.updated == true {
		t := suite.T()
		t.Run("delete classroom", func(t *testing.T) {
			classroom, err := entity.NewClassroom(
				"Other Name",
				"Other Description",
				suite.course,
			)
			if err != nil {
				panic(err.Error())
			}

			err = suite.repository.Create(classroom)
			assert.NoError(t, err)
			assert.NotNil(t, classroom.ID)

			id, err := entity.NewInputID(classroom.ID)
			assert.NoError(t, err)

			err = suite.repository.Delete(id)
			assert.NoError(t, err)

			currentClassroom, err := suite.repository.Show(id)
			assert.Error(t, err)
			assert.Nil(t, currentClassroom)
		})
	}
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
	suite.created = true
}

func (suite *ClassroomRepoTestSuite) TestShowClassroom() {
	t := suite.T()

	currentClassroom, err := suite.repository.Show(suite.classroom)
	assert.NoError(t, err)
	assert.NotNil(t, currentClassroom)
	assert.Equal(t, "Name", currentClassroom.Name)
	assert.Equal(t, "Description", currentClassroom.Description)
	assert.Equal(t, suite.course, currentClassroom.Course)
	suite.showed = true
}

func (suite *ClassroomRepoTestSuite) TestUpdateClassroom() {
	t := suite.T()

	newClassroom, err := entity.UpdateClassroom(
		"New Name",
		"New Description",
		suite.course,
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.repository.Update(suite.classroom, newClassroom)
	assert.NoError(t, err)
	assert.NotNil(t, newClassroom)
	assert.NotEqual(t, "Name", newClassroom.Name)
	assert.NotEqual(t, "Description", newClassroom.Description)
	assert.Equal(t, suite.course, newClassroom.Course)
	suite.updated = true
}

func TestClassroomRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ClassroomRepoTestSuite))
}
