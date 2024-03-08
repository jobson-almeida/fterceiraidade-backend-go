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

type AssessmentRepoTestSuite struct {
	suite.Suite
	container           *testhelpers.DatabaseContainer
	ctx                 context.Context
	conn                *gorm.DB
	courseRepository    *repository.CourseRepository
	classroomRepository *repository.ClassroomRepository
	course              string

	// assessment *entity.InputID
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
		})
	}
}

func (suite *AssessmentRepoTestSuite) AfterTest(_, _ string) {
}

func (suite *AssessmentRepoTestSuite) TestCreateAssessment() {
}

func TestAssessmentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AssessmentRepoTestSuite))
}
