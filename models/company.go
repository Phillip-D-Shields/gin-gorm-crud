package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID                 uint   `json:"id" gorm:"primary_key"`
	Name               string `json:"name" gorm:"unique"`
	BillingAddress     string `json:"billing_address"`
	ShippingAddress    string `json:"shipping_address"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	PrimaryContactID   int    `json:"primary_contact_id,omitempty"`
	SecondaryContactID int    `json:"secondary_contact_id,omitempty"`
	Notes              string `json:"notes"`
}
