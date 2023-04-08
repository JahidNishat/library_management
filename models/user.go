package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `gorm:"unique" json:"email"`
	Phone        string `gorm:"unique" json:"phone"`
	Password     string `json:"password"`
	UserType     string `json:"user_type"`
	UserId       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
