package api

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/controller"
	"github.com/library_management/routes"
)

func StartServer() {
	router := gin.Default()
	router.Use(gin.Logger())

	router.POST("/book/", controller.CreateBook)
	router.PUT("/book/:bookId", controller.UpdateBookById)
	router.DELETE("/book/:bookId", controller.DeleteBookById)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run()
}
