package models

import (
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model         // This will include ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint    `json:"user_id"` // Foreign Key for User
	User       User    `json:"user"`    // The User this budget belongs to
	Amount     float64 `json:"amount"`
	Month      int     `json:"month"`
	Year       int     `json:"year"`
}
