package models

import "gorm.io/gorm"

//User model contains first,Last name and email address
type User struct {
	gorm.Model
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Post      []Posts    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one user can have many posts
	PostLike  []PostLike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one user like many post
	Role      string     `json:"role"`
}

//Not a model only aa stuct
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
