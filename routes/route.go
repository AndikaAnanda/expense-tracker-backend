package routes

import (
	"expense-tracker-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

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