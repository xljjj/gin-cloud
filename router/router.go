package router

import (
	"CloudDrive/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("", controller.Index)
	router.GET("/index", controller.Index)

	return router
}
