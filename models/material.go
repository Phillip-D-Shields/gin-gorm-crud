package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Code        string  `json:"material_code" gorm:"unique;not null"`
	Description string  `json:"description"`
	CostPerUnit float64 `json:"cost_per_unit"`
	SupplierID  uint    `json:"supplier_id"`
	Notes       string  `json:"notes"`
}
