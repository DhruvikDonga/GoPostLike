package models

import "gorm.io/gorm"

//User model contains first,Last name and email address
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
