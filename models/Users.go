package models

import "gorm.io/gorm"

//User model contains first,Last name and email address
type User struct {
	gorm.Model
	Username       string      `json:"username"`
	FirstName      string      `json:"firstname"`
	LastName       string      `json:"lastname"`
	Email          string      `json:"email"`
	Password       string      `json:"password"`
	Post           []Posts     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one user can have many posts
	PostLike       []PostLike  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one user like many post
	Role           string      `json:"role"`
	TotalPosts     int         `json:"totalposts"`
	TotalFollowers int         `json:"totalfollowers"`
	Followers      []Followers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one user can have many followers
}

//Followers model is many to many relationship between usersid
type Followers struct {
	gorm.Model
	UserID     int `json:"userid"`
	FollowerID int `json:"followerid"`
}

//Not a model only a login structure
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
