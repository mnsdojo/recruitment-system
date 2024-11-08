package models

import "time"

type Job struct {
	PostedOn          time.Time `json:"posted_on"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	CompanyName       string    `json:"company_name"`
	ID                int       `json:"id" gorm:"primaryKey"`
	TotalApplications int       `json:"total_applications"`
	PostedBy          int       `json:"posted_by"`
}
