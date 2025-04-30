package database

import (
	"fmt"
	"log"

	"github.com/quantinium3/lucy/cmd/database/model"
	"github.com/quantinium3/lucy/cmd/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func ConnectDB() (*Database, error) {
	dbPort := utils.Config("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT environment variable in undefined")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", utils.Config("DB_HOST"), utils.Config("DB_USER"), utils.Config("DB_PASSWORD"), utils.Config("DB_NAME"), dbPort, utils.Config("DB_SSLMODE"), utils.Config("DB_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Database Connected")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&model.User{}, &model.Peripheral{})

	return &Database{DB: db}, nil
}
