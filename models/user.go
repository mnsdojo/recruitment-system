// models/user.go
package models

import "gorm.io/gorm"

type UserType string

const (
	Admin     UserType = "Admin"
	Applicant UserType = "Applicant"
)

type User struct {
	gorm.Model
	Name            string   `json:"name" gorm:"not null"`
	Email           string   `json:"email" gorm:"unique;not null"`
	PasswordHash    string   `json:"-"` // password hash, not returned in JSON response
	UserType        UserType `json:"user_type" gorm:"not null"`
	ProfileHeadline string   `json:"profile_headline"`
	Address         string   `json:"address"`
	Profile         Profile  `json:"profile" gorm:"foreignKey:ApplicantID"`
}
