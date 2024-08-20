package controller

import (
	"CloudDrive/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Help(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUser(fmt.Sprintf("%v", openId))

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)

	c.HTML(http.StatusOK, "help.html", gin.H{
		"currHelp":      "active",
		"user":          user,
		"fileDetailUse": fileDetailUse,
	})
}
