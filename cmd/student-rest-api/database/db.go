package database

import (
	"fmt"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file, error=%s", err.Error())
	}

	//Data Source Name-> dsn, connStr
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Error connecting to database, %s", err.Error())
	}

	log.Println("Successfully connected to the DB")
	DB = db

	//Auto Migrate
	if err := db.AutoMigrate(&models.Student{}, &models.Course{}, &models.Teacher{}, &models.Users{}); err != nil {
		return err
	}

	return nil
}
