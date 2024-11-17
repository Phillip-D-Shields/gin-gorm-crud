package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone"`
	Notes     string `json:"notes,omitempty"`
}
