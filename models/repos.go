package models

import (
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetAllBooks() []Book{
	var books []Book
	Db.Find(&books)
	return books
}

func (b *Book) CreateBook() *Book{
	Db.Create(&b)
	return b
}