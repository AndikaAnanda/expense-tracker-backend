package main

import (
	"log"
	"os"

	"expense-tracker-backend/config"
	"expense-tracker-backend/models"
	"expense-tracker-backend/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Schema migration
	if err := config.DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Printf("Migration warning: %v", err) // jangan Fatalf, supaya app tetap jalan
	} else {
		log.Println("Migration success")
	}

	app := routes.SetupRoutes()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}