package router

import (
	"CloudDrive/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("", controller.Welcome)
	router.GET("/welcome", controller.Welcome)
	router.GET("/login", controller.Login)
	router.GET("/qq_login", controller.HandleLogin)

	return router
}
