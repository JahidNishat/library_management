package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"uniqueIndex"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
