package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // This will include ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	Currency   string    `json:"currency"`
	Expenses   []Expense `json:"expenses"` // One-to-many relationship
	Budgets    []Budget  `json:"budgets"`  // One-to-many relationship
}
