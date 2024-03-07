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
	showed     bool
	updated    bool
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
	if suite.created == true && suite.showed == true && suite.updated == true {
		t := suite.T()
		t.Run("delete student", func(t *testing.T) {
			id, err := entity.NewInputID(suite.student.ID)
			assert.NoError(t, err)

			err = suite.repository.Delete(id)
			assert.NoError(t, err)

			currentStudent, err := suite.repository.Show(id)
			assert.Error(t, err)
			assert.Nil(t, currentStudent)
		})
	}
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

func (suite *StudentRepoTestSuite) TestShowStudent() {
	t := suite.T()

	currentStudent, err := suite.repository.Show(suite.student)
	assert.NoError(t, err)
	assert.NotNil(t, currentStudent)
	assert.Equal(t, "/avatar/avatar.png", currentStudent.Avatar)
	assert.Equal(t, "Firstname", currentStudent.Firstname)
	assert.Equal(t, "Lastname", currentStudent.Lastname)
	assert.Equal(t, "email@email.com", currentStudent.Email)
	assert.Equal(t, "+998812345678", currentStudent.Phone)
	assert.Equal(t, "City", currentStudent.Address.City)
	assert.Equal(t, "State", currentStudent.Address.State)
	assert.Equal(t, "Street", currentStudent.Address.Street)
	suite.showed = true
}

func (suite *StudentRepoTestSuite) TestUpdateStudent() {
	t := suite.T()

	newAddress := entity.DetailsAddress{
		City: "New City", State: "New State", Street: "New Street",
	}

	updateStudent, err := entity.UpdateStudent(
		"/avatar/new_avatar.png",
		"New Firstname",
		"New Lastname",
		"new_email@email.com",
		"+008812345678",
		newAddress,
	)
	if err != nil {
		panic(err.Error())
	}

	err = suite.repository.Update(suite.student, updateStudent)
	assert.NoError(t, err)
	assert.Equal(t, "/avatar/new_avatar.png", updateStudent.Avatar)
	assert.Equal(t, "New Firstname", updateStudent.Firstname)
	assert.Equal(t, "New Lastname", updateStudent.Lastname)
	assert.Equal(t, "new_email@email.com", updateStudent.Email)
	assert.Equal(t, "+008812345678", updateStudent.Phone)
	assert.Equal(t, "New City", updateStudent.Address.City)
	assert.Equal(t, "New State", updateStudent.Address.State)
	assert.Equal(t, "New Street", updateStudent.Address.Street)
	suite.updated = true
}

func TestStudentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepoTestSuite))
}
