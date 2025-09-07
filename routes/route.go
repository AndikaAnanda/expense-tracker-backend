package routes

import (
	"github.com/gin-gonic/gin"
	"expense-tracker-backend/controllers"
)

func SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	{
		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("/transactions/:id", controllers.DeleteTransaction)
	}
}
