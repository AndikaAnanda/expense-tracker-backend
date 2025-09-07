package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"expense-tracker-backend/config"
	"expense-tracker-backend/models"
)

type CreateTransactionDTO struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Type   string  `json:"type" binding:"required,oneof=income expense"`
}

type UpdateTransactionDTO struct {
	Title  *string  `json:"title,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Type   *string  `json:"type,omitempty" binding:"omitempty,oneof=income expense"`
}

// GET /transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	result := config.DB.Order("created_at_desc").Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// POST /transactions
func CreateTransaction(c *gin.Context) {
	var dto CreateTransactionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := models.Transaction{
		Title: dto.Title,
		Amount: dto.Amount,
		Type: dto.Type,
	}

	if err := config.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

// PUT /transactions/:id
func UpdateTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tx models.Transaction
	if err := config.DB.First(&tx, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	var dto UpdateTransactionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dto.Title != nil {
		tx.Title = *dto.Title
	}

	if dto.Amount != nil {
		tx.Amount = *dto.Amount
	}
	if dto.Type != nil {
		tx.Type = *dto.Type
	}

	if err := config.DB.Save(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

// DELETE /transactions/:id
func DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tx models.Transaction
	if err := config.DB.First(&tx, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	if err := config.DB.Delete(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}

// GET /transactions/summary (count total income, expense, and balance)
func GetSummary(c *gin.Context) {
	type Row struct {Total float64}
	var inc, exp Row

	config.DB.Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount),0) as total").
		Where("type = ?", "income").Scan(&inc)

	config.DB.Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount),0) as total").
		Where("type = ?", "expense").Scan(&exp)

	c.JSON(http.StatusOK, gin.H{
		"total income": inc.Total,
		"total expense": exp.Total,
		"total balance": inc.Total - exp.Total,
	})
}
