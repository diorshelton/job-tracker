package models

type CompanyOfInterest struct {
	ID int `json:"id"`
	CompanyName string `json:"company_name"`
	Location    string `json:"location,omitempty"`
	Industry    string `json:"industry,omitempty"`
	Technologies string `json:"technologies,omitempty"`
	URL         string `json:"url,omitempty"`
	Notes       string `json:"notes,omitempty"`
}