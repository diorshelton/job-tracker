package models
import (
	"time"
)

type Job struct {
	ID         int       `json:"id"`
	Company    string    `json:"company"`
	Title      string    `json:"title"`
	Location   string    `json:"location,omitempty"`
	URL        string    `json:"url,omitempty"`
	DatePosted time.Time `json:"date_posted,omitempty"`
}