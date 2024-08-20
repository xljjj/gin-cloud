package router

import (
	"CloudDrive/controller"
	"CloudDrive/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("", controller.Welcome)
	router.GET("/welcome", controller.Welcome)
	router.GET("/login", controller.Login)
	router.GET("/qq_login", controller.HandleLogin)

	cloud := router.Group("cloud")
	cloud.Use(middleware.CheckLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/help", controller.Help)
	}
	return router
}
