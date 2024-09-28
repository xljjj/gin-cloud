package router

import (
	"CloudDrive/controller"
	"CloudDrive/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CheckInput)

	router.GET("", controller.Welcome)
	router.GET("/welcome", controller.Welcome)
	router.GET("/login", controller.Login)
	router.GET("/qq-login", controller.HandleLogin)
	router.GET("/register", controller.Register)
	router.GET("/help", controller.Help)

	router.POST("/register", controller.HandleRegister)
	router.POST("/simple-login", controller.HandleSimpleLogin)

	// 自定义404页面
	router.NoRoute(func(c *gin.Context) {
		// 返回自定义的404页面内容
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	cloud := router.Group("cloud")
	cloud.Use(middleware.CheckSimpleLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/logout", controller.SimpleLogout)
		cloud.GET("/file", controller.File)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/downloadFile", controller.DownloadFile)
		cloud.GET("/viewFile", controller.ViewFile)
		cloud.GET("/modify", controller.Modify)
	}
	{
		cloud.POST("/addFolder", controller.AddFolder)
		cloud.POST("/updateFolder", controller.UpdateFolder)
		cloud.POST("/uploadFile", controller.HandleUpload)
		cloud.POST("/getQRCode", controller.ShareFile)
		cloud.POST("/modify", controller.HandleModify)
	}
	{
		cloud.DELETE("/deleteFile", controller.DeleteFile)
		cloud.DELETE("/deleteFolder", controller.DeleteFolder)
	}

	admin := router.Group("admin")
	admin.Use(middleware.CheckAdmin)
	{
		admin.GET("/index", controller.Admin)
	}
	{
		admin.POST("/query", controller.QuerySimpleUser)
		admin.POST("/delete", controller.DeleteSimpleUser)
	}

	return router
}
