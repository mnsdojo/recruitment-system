// models/profile.go
package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ResumeFilePath string `json:"resume_file_path"`
	Skills         string `json:"skills"`
	Education      string `json:"education"`
	Experience     string `json:"experience"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	ApplicantID    uint   `json:"applicant_id" gorm:"not null;unique"`
}
