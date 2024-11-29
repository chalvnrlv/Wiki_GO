package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model            // This will include ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Expenses    []Expense `json:"expenses"` // One-to-many relationship
}
