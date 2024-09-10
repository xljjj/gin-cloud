package controller

import (
	"CloudDrive/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Admin 管理员页面
func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"hint": "欢迎登录管理员页面",
	})
}

// QuerySimpleUser 查询用户信息
func QuerySimpleUser(c *gin.Context) {
	username := c.PostForm("username")
	s := model.FindSimpleUserByUserName(username)
	if s.UserName == "" {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"hint": "该用户名不存在！",
		})
	} else {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"hint": "用户查询成功！",
			"user": s,
		})
	}
}

// DeleteSimpleUser 删除用户
func DeleteSimpleUser(c *gin.Context) {
	username := c.PostForm("username")
	s := model.FindSimpleUserByUserName(username)
	if s.UserName == "" {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"hint": "该用户名不存在！",
		})
	} else {
		if s.Ext != "" {
			err := os.Remove("./avatar/" + s.UserName + s.Ext)
			if err != nil {
				c.HTML(http.StatusOK, "admin.html", gin.H{
					"hint": "删除用户失败！",
				})
				return
			}
		}
		err := model.DeleteFileStore(s.FileStoreId)
		if err != nil {
			c.HTML(http.StatusOK, "admin.html", gin.H{
				"hint": "删除用户失败！",
			})
			return
		}
		model.DeleteSimpleUser(&s)
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"hint": "用户删除成功！",
		})
	}
}
