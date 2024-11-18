package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model        // includes ID, CreatedAt, UpdatedAt, DeletedAt
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email" gorm:"uniqueIndex"`
	Phone      string `json:"phone"`
	Notes      string `json:"notes,omitempty"`
}
