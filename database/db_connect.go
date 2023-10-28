package database

import (
	"fmt"
	"log"
	"os"

	"github.com/amirthapa27/precize-test/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	// connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	// connect to postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// check for error
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	// migrate model
	err = DB.AutoMigrate(models.SAT_Results{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database migrated successfully")
}
