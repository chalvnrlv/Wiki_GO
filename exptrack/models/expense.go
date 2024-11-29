package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model           // This will include ID, CreatedAt, UpdatedAt, DeletedAt
	UserID      uint     `json:"user_id"` // Foreign Key for User
	User        User     `json:"user"`    // The User this expense belongs to
	Amount      float64  `json:"amount"`
	CategoryID  uint     `json:"category_id"` // Foreign Key for Category
	Category    Category `json:"category"`    // The Category this expense belongs to
	Date        string   `json:"date"`
	Description string   `json:"description"`
}
