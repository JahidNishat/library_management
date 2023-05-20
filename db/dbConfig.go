package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/library_management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/go-gormigrate/gormigrate/v2"
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

	//Migration
	m := gormigrate.New(Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230512000001",
			Migrate: func(tx *gorm.DB) error {
				// Create new table or perform any other migration operations
				return tx.AutoMigrate(&models.Token{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback migration (if needed)
				return tx.Migrator().DropTable(&models.Token{})
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalln(err)
	}

	// Db.AutoMigrate(&models.Book{}, &models.User{}, &models.Token{})

	return Db
}