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
	router.GET("/qq-login", controller.HandleLogin)
	router.GET("/register", controller.Register)

	router.POST("/register", controller.HandleRegister)
	router.POST("/simple-login", controller.HandleSimpleLogin)

	cloud := router.Group("cloud")
	cloud.Use(middleware.CheckSimpleLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/help", controller.Help)
		cloud.GET("/file", controller.File)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/logout", controller.SimpleLogout)
		cloud.GET("/downloadFile", controller.DownloadFile)
		cloud.GET("/deleteFile", controller.DeleteFile)
		cloud.GET("/deleteFolder", controller.DeleteFileFolder)
		cloud.GET("/modify", controller.Modify)
	}
	{
		cloud.POST("/addFolder", controller.AddFolder)
		cloud.POST("/updateFolder", controller.UpdateFileFolder)
		cloud.POST("/uploadFile", controller.HandleUpload)
		cloud.POST("/getQRCode", controller.ShareFile)
		cloud.POST("/modify", controller.HandleModify)
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
