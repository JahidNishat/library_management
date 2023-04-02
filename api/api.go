package api

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/config"
	"github.com/library_management/controller"
)

func StartServer() {
	Router := gin.Default()
	config.Connect()

	Router.GET("/book/", controller.GetBook)
	Router.POST("/book/", controller.CreateBook)
	// Router.Get("/book/:bookId", GetBookById)
	// Router.Put("/book/:bookId", UpdateBookById)
	// Router.Delete("/book/:bookId", DeleteBookById)

	Router.Run()
}