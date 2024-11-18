package models

import "gorm.io/gorm"

// TODO - is this the best way to handle status?
type Status int

const (
	JobStatusBacklog Status = iota
	JobStatusPending
	JobStatusInProgress
	JobStatusCompleted
	JobStatusInvoiced
	JobStatusPaid
	JobStatusCancelled
)

var statusText = map[Status]string{
	JobStatusBacklog:    "backlog",
	JobStatusPending:    "pending",
	JobStatusInProgress: "in_progress",
	JobStatusCompleted:  "completed",
	JobStatusInvoiced:   "invoiced",
	JobStatusPaid:       "paid",
	JobStatusCancelled:  "cancelled",
}

func (s Status) String() string {
	return statusText[s]
}

type Job struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PriceQuote  float64    `json:"price_quote"`
	CompanyID   uint       `json:"company_id,omitempty"`
	ContactID   uint       `json:"contact_id,omitempty"`
	Services    []Service  `json:"services" gorm:"many2many:job_services;"`
	Materials   []Material `json:"materials" gorm:"many2many:job_materials;"`
	Status      Status     `json:"status"`
	Notes       string     `json:"notes"`
}
