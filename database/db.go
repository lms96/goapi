package database

import (
	"goapi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	connection := "host=localhost user=postgres password=q1w2e3r4t5 dbname=goapi port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))

	if err != nil {
		log.Panic("Error connecting to database.")
	}

	DB.AutoMigrate(&models.Student{})
}
