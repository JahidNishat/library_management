package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/controller"
	"github.com/library_management/middleware"
)

func AuthRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate)

	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:user_id", controller.GetUserById)
	router.DELETE("/users/:user_id", controller.DeleteUserById)
	router.GET("/users/logout", controller.LogOut)
	router.GET("/users/refresh", controller.RefreshToken)

	router.GET("/book/", controller.GetAllBooks)
	router.GET("/book/:bookId", controller.GetBookById)
	router.POST("/book/", controller.CreateBook)
	router.PUT("/book/:bookId", controller.UpdateBookById)
	router.DELETE("/book/:bookId", controller.DeleteBookById)
}
