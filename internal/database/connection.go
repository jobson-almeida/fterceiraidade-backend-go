package database

import (
	"fmt"
	"os"

	entity "github.com/jobson-almeida/fterceiraidade-backend-go/internal/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dialect := os.Getenv("DIALECT")
	connection := os.Getenv("CONNECTION")

	conn, err := gorm.Open(postgres.New(postgres.Config{DriverName: dialect, DSN: connection}))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection successful")
	conn.AutoMigrate(
		entity.Course{},
		entity.Student{},
		entity.Teacher{},
		entity.Assessment{},
		entity.Classroom{},
		entity.Question{},
	)

	return conn
}
