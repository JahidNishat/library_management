package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/controller"
	"github.com/library_management/middleware"
)

func AuthRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate)

	router.GET("/users")
	router.GET("/users/:user_id")

	router.POST("/book/", controller.CreateBook)
	router.PUT("/book/:bookId", controller.UpdateBookById)
	router.DELETE("/book/:bookId", controller.DeleteBookById)
}
