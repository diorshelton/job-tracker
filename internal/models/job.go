package models
import (
	"time"
)

type Job struct {
	ID          int               `json:"id"`
	JobID       int               `json:"job_id"`
	Company string            `json:"company"`
	Position string            `json:"position"`
	Status      ApplicationStatus `json:"status"`
	Location string            `json:"location,omitempty"`
	URL         string            `json:"url,omitempty"`
	Notes       string            `json:"notes,omitempty"`
	DatePosted time.Time         `json:"date_posted,omitempty"`
	DateApplied time.Time         `json:"date_applied"`
	LastUpdated time.Time         `json:"last_updated"`
}

type ApplicationStatus string

const (
	StatusApplied    ApplicationStatus = "applied"
	StatusScreening  ApplicationStatus = "screening"
	StatusAssessment ApplicationStatus = "assessment"
	StatusInterview  ApplicationStatus = "interview"
	StatusRejected   ApplicationStatus = "rejected"
	StatusOffer      ApplicationStatus = "offer"
	StatusArchived   ApplicationStatus = "archived"
)



