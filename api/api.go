package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/library_management/routes"
	"github.com/spf13/viper"
)

func StartServer() {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}

	router := gin.Default()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	port := viper.GetString("server.port")
	router.Run(":" + port)
}
