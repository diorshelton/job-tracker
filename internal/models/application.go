package models

import (
	"time"
)

type ApplicationStatus string

type Application struct {
	ID          int               `json:"id"`
	JobID       int               `json:"job_id"`
	Status      ApplicationStatus `json:"status"`
	DateApplied time.Time         `json:"date_applied"`
	LastUpdated time.Time         `json:"last_updated"`
	Notes       string            `json:"notes,omitempty"`
}

const (
	StatusApplied    ApplicationStatus = "applied"
	StatusScreening  ApplicationStatus = "screening"
	StatusAssessment ApplicationStatus = "assessment"
	StatusInterview  ApplicationStatus = "interview"
	StatusRejected   ApplicationStatus = "rejected"
	StatusOffer      ApplicationStatus = "offer"
	StatusArchived   ApplicationStatus = "archived"
)


