package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library_management/models"
)

func GetBook(ctx *gin.Context){
	books := models.GetAllBooks()
	ctx.JSON(200, gin.H{
		"data": books,
	})
}

func CreateBook(ctx *gin.Context){
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	b := book.CreateBook()
	ctx.JSON(200, gin.H{
		"data": b,
	})
}