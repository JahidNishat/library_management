package db

import (
	"fmt"
	"github.com/spf13/viper"
	"log"

	"github.com/library_management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	Dbname   string
}

func Connect() *gorm.DB {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}

	config := &Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Password: viper.GetString("database.password"),
		User:     viper.GetString("database.user"),
		Dbname:   viper.GetString("database.dbname"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.Dbname, config.Port,
	)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error", err)
	}

	Db.AutoMigrate(&models.Book{}, &models.User{})

	return Db
}
