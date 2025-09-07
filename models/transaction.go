package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Type      string    `json:"type" binding:"required,oneof=income expense"` // income / expense
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}