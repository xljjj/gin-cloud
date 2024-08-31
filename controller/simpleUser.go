package controller

import (
	"CloudDrive/model"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// Register 注册用户
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"hint": "",
	})
}

// HandleRegister 处理注册请求
func HandleRegister(c *gin.Context) {
	// 获取表单中的数据
	username := c.PostForm("username")               // 获取用户名
	password := c.PostForm("password")               // 获取密码
	confirmPassword := c.PostForm("confirmPassword") //获取确认密码
	nickname := c.PostForm("nickname")               // 获取昵称
	avatar, _ := c.FormFile("avatar")                //获取头像

	var hint string

	// 验证用户名是否为 8-30 位
	if len(username) < 5 || len(username) > 30 {
		hint = "用户名必须为  到 30 位！"
	}

	// 验证密码和确认密码是否匹配
	if password != confirmPassword {
		hint = "密码和确认密码不匹配！"
	}

	// 验证昵称是否为 1-10 位
	if len(nickname) < 1 || len(nickname) > 10 {
		hint = "昵称必须为 1 到 10 位！"
	}

	// 验证头像文件格式
	if avatar != nil {
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/gif":  true,
		}
		if !allowedTypes[avatar.Header.Get("Content-Type")] {
			hint = "头像文件格式不正确，请上传 JPEG、PNG 或 GIF 图片！"
		}
	}

	// 验证用户名是否已存在
	s := model.FindSimpleUserByUserName(username)
	if s.UserName != "" {
		hint = "该用户名已存在！"
	}

	if hint != "" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"hint": hint,
		})
		return
	}

	// 成功注册逻辑
	if avatar != nil {
		_ = c.SaveUploadedFile(avatar, "./avatar/"+username+filepath.Ext(avatar.Filename))
	} else {
		fmt.Println("没有头像保存")
	}

	user := model.SimpleUser{
		UserName: username,
		Password: util.Md5Encode(password),
		NickName: nickname,
	}
	if avatar != nil {
		user.Ext = filepath.Ext(avatar.Filename)
	}
	model.CreateSimpleUser(&user)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"status": "success",
		"hint":   "注册成功,欢迎登录",
	})
}
