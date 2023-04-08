package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/controller"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.LogIn)

	router.GET("/book/", controller.GetAllBooks)
	router.GET("/book/:bookId", controller.GetBookById)
}