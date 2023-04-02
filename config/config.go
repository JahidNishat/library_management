package config

import (
	"log"

	"github.com/library_management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB{
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error", err)
	}

	Db.AutoMigrate(&models.Book{})

	return Db
}