package entity_test

import (
	"testing"

	"github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"

	"github.com/stretchr/testify/require"
)

func TestNewStudent(t *testing.T) {
	address := entity.DetailsAddress{
		City:   "city",
		State:  "state",
		Street: "street",
	}
	_, err := entity.NewStudent("avatar", "firstname", "lastname", "email@email.com", "+5512345678", address)

	require.Nil(t, err)
	// require.Error(t, err)
}

/*
func TestNewStudentf(t *testing.T) {
	student := entity.NewStudent()
	student.Base = entity.Base{
		ID:        uuid.New().String(),
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	//DeletedAt:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},

	student.Avatar = "Avatar"
	student.Firstname = "Firstname"
	student.Lastname = "Lastname"
	student.Email = "emailfemail.com"
	student.Phone = "+5512345678"
	student.Address = entity.DetailsAddress{City: "City", State: "State", Street: "Street"}

	require.NotNil(t, student)
}
*/
