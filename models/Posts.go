package models

import "gorm.io/gorm"

//Posts model contains details of posts by users
type Posts struct {
	*gorm.Model
	Postname        string      `json:"postname"`
	Postslug        string      `json:"postslug"`
	Postdescription string      `json:"description"`
	PostLikes       int         `json:"likes"`
	UserID          int         `json:"userid"`
	PostImage       []PostImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one post can have many images
	PostLike        []PostLike  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //one post can have many likes
}

//PostImage contains multiple images of a post
type PostImage struct {
	*gorm.Model
	PostsID int    `json:"postid"`
	Image   string `json:"image"`
}

//PostLike Many-to-Many relationship between user likes the posts
type PostLike struct {
	*gorm.Model
	PostsID int `json:"postid"`
	UserID  int `json:"userid"`
}
