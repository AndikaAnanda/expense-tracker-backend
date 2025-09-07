package config

import (
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	// initialize .env
	_= godotenv.Load()
}

func ConnectDB() {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")
	// name := os.Getenv("DB_NAME")
	// sslmode := os.Getenv("DB_SSLMODE")
	// tz := os.Getenv("DB_TIMEZONE")

	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	// 	host, user, pass, name, port, sslmode, tz,
	// )

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	DB = db
	log.Println("Database connected")
}