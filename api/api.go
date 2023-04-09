package api

import (
	"github.com/gin-gonic/gin"
	"github.com/library_management/routes"
)

func StartServer() {
	router := gin.Default()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	
	router.Run()
}
