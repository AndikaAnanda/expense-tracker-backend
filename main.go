package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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

	app := gin.Default()
		
	// CORS setup
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://your-frontend-domain.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes()
	
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}