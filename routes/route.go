package routes

import (
	"expense-tracker-backend/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// CORS setup
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }))

	api := r.Group("/api") 
	{
		api.GET("/health", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"status": "ok"})})	
	
		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("transactions/:id", controllers.DeleteTransaction)
	
		api.GET("/transactions/summary", controllers.GetSummary)
	}

	return r
}