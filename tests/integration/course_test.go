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

type CourseRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.CourseRepository
	ctx        context.Context
	course     *entity.InputID
	// created    bool
	// showed     bool
	// updated    bool
}

func (suite *CourseRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	container, err := testhelpers.CreateDatabaseContainer(suite.ctx, "courses.sql")
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
	repository := repository.NewCourseRepository(conn)

	conn.AutoMigrate(
		entity.Course{},
	)
	suite.repository = repository
}

func (suite *CourseRepoTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *CourseRepoTestSuite) AfterTest(_, _ string) {
}

func (suite *CourseRepoTestSuite) TestCreateCourse() {
	t := suite.T()

	course, err := entity.NewCourse(
		"Name",
		"Description",
		"/image/image.png",
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.repository.Create(course)
	assert.NoError(t, err)
	assert.NotNil(t, course)

	suite.course, err = entity.NewInputID(course.ID)
	assert.NoError(t, err)
}

func (suite *CourseRepoTestSuite) TestShowCourse() {
	t := suite.T()

	currentCourse, err := suite.repository.Show(suite.course)
	assert.NoError(t, err)
	assert.NotNil(t, currentCourse)
	assert.Equal(t, "Name", currentCourse.Name)
	assert.Equal(t, "Description", currentCourse.Description)
	assert.Equal(t, "/image/image.png", currentCourse.Image)
}

func TestCourseRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRepoTestSuite))
}
