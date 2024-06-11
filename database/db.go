package db

import (
	"log"

	student "github.com/Vinimolz/Gin_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func DatabaseConnection() {
	connectionString := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Error connecting to the database")
	}

	Db.AutoMigrate(&student.Student{})
}
