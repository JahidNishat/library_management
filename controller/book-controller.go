package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library_management/config"
	"github.com/library_management/helper"
	"github.com/library_management/models"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db = config.Connect()
}

func GetAllBooks(ctx *gin.Context) {
	var books []models.Book
	res := Db.Find(&books)
	if res.Error != nil {
		log.Fatalln("DB Retrieve Data Error")
	}

	ctx.JSON(200, gin.H{
		"data": books,
	})
}

func GetBookById(ctx *gin.Context) {
	id := ctx.Param("bookId")
	var book models.Book
	if err := Db.Where(id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Record Not Found")
		return
	}

	ctx.JSON(200, gin.H{
		"data": book,
	})
}

func CreateBook(ctx *gin.Context) {
	if err := helper.CheckUserType(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := ctx.Bind(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Db.Create(&book)
	ctx.JSON(200, gin.H{
		"data": book,
	})
}

func UpdateBookById(ctx *gin.Context) {
	if err := helper.CheckUserType(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("bookId")
	var book models.Book
	if err := Db.Where(id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Record Not Found")
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Db.Updates(book)

	ctx.JSON(200, gin.H{
		"data": book,
	})
}

func DeleteBookById(ctx *gin.Context) {
	if err := helper.CheckUserType(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("bookId")
	var book models.Book
	if err := Db.Where(id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Record Not Found")
		return
	}

	Db.Unscoped().Delete(&book)
	ctx.JSON(200, gin.H{
		"data": book,
	})
}
