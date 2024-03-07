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
	"github.com/stretchr/testify/suite"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClassroomRepoTestSuite struct {
	suite.Suite
	container  *testhelpers.DatabaseContainer
	repository *repository.ClassroomRepository
	ctx        context.Context
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

func TestClassroomRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ClassroomRepoTestSuite))
}
