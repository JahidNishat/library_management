package api

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/config"
	"github.com/library_management/controller"
)

func StartServer() {
	Router := gin.Default()
	config.Connect()

	Router.GET("/book/", controller.GetAllBooks)
	Router.POST("/book/", controller.CreateBook)
	Router.GET("/book/:bookId", controller.GetBookById)
	Router.PUT("/book/:bookId", controller.UpdateBookById)
	Router.DELETE("/book/:bookId", controller.DeleteBookById)

	Router.Run()
}