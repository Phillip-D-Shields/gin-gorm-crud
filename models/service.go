package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Code        string     `json:"service_code" gorm:"unique;not null"`
	Description string     `json:"description"`
	RatePerHour float64    `json:"rate_per_hour"`
	RatePerItem float64    `json:"rate_per_item"`
	Materials   []Material `json:"materials" gorm:"many2many:service_materials;"`
	Notes       string     `json:"notes"`
}
