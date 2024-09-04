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
	router.GET("/register", controller.Register)
	router.GET("/admin", controller.Admin)
	router.GET("/modify", controller.Modify)

	router.POST("/register", controller.HandleRegister)
	router.POST("/query", controller.QuerySimpleUser)
	router.POST("/delete", controller.DeleteSimpleUser)
	router.POST("/modify", controller.HandleModify)

	cloud := router.Group("cloud")
	cloud.Use(middleware.CheckLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/help", controller.Help)
		cloud.GET("/file", controller.File)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/logout", controller.Logout)
		cloud.GET("/downloadFile", controller.DownloadFile)
		cloud.GET("/deleteFile", controller.DeleteFile)
		cloud.GET("/deleteFolder", controller.DeleteFileFolder)
		cloud.GET("/doc-file", controller.DocFile)
		cloud.GET("/image-file", controller.ImageFile)
		cloud.GET("/video-file", controller.VideoFile)
		cloud.GET("/music-file", controller.MusicFile)
		cloud.GET("/other-file", controller.OtherFile)
	}
	{
		cloud.POST("/addFolder", controller.AddFolder)
		cloud.POST("/updateFolder", controller.UpdateFileFolder)
		cloud.POST("/uploadFile", controller.HandleUpload)
		cloud.POST("/getQRCode", controller.ShareFile)
	}
	return router
}
