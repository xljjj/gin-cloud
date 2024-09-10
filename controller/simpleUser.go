package controller

import (
	"CloudDrive/model"
	"CloudDrive/mysql"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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
	userName := c.PostForm("username")               // 获取用户名
	password := c.PostForm("password")               // 获取密码
	confirmPassword := c.PostForm("confirmPassword") //获取确认密码
	nickName := c.PostForm("nickname")               // 获取昵称
	avatar, _ := c.FormFile("avatar")                //获取头像

	var hint string

	// 验证用户名是否为 8-30 位
	if len(userName) < 5 || len(userName) > 30 {
		hint = "用户名必须为 5 到 30 位！"
	}

	// 验证密码和确认密码是否匹配
	if password != confirmPassword {
		hint = "密码和确认密码不匹配！"
	}

	// 验证昵称是否为 1-10 位
	if len(nickName) < 1 || len(nickName) > 10 {
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
	s := model.FindSimpleUserByUserName(userName)
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
		_ = c.SaveUploadedFile(avatar, "./avatar/"+userName+filepath.Ext(avatar.Filename))
	}

	user := model.SimpleUser{
		UserName: userName,
		Password: util.Md5Encode(password),
		NickName: nickName,
	}
	if avatar != nil {
		user.Ext = filepath.Ext(avatar.Filename)
	}
	model.CreateSimpleUser(&user)
	// 创建文件仓库
	fileStore := model.FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	mysql.DB.Create(&fileStore)
	user.FileStoreId = fileStore.Id
	mysql.DB.Save(&user)
	// 创建一个属于用户的文件夹（代指文件仓库）
	_ = os.Mkdir(strconv.Itoa(fileStore.Id), os.ModePerm)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"status": "success",
		"hint":   "注册成功,欢迎登录",
	})
}

// Modify 修改用户信息
func Modify(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	user := model.FindSimpleUserByUserName(userName)
	c.HTML(http.StatusOK, "modify.html", gin.H{
		"hint": "",
		"user": user,
	})
}

// HandleModify 处理修改请求
func HandleModify(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	currentPassword := c.PostForm("currentPassword")
	newPassword := c.PostForm("newPassword")
	confirmPassword := c.PostForm("confirmPassword")
	nickName := c.PostForm("nickname")
	avatar, _ := c.FormFile("avatar")

	var hint string

	// 验证密码和确认密码是否匹配
	if newPassword != confirmPassword {
		hint = "新密码和确认密码不匹配！"
	}

	// 验证昵称是否为 1-10 位
	if len(nickName) > 10 {
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

	// 检查用户原账户
	s := model.FindSimpleUserByUserName(userName)
	if s.UserName == "" {
		hint = "该用户不存在！"
	}
	if util.Md5Encode(currentPassword) != s.Password {
		hint = "原密码错误！"
	}

	if hint != "" {
		c.HTML(http.StatusOK, "modify.html", gin.H{
			"hint": hint,
		})
		return
	}

	// 成功逻辑
	if avatar != nil {
		// 删除原头像文件
		if s.Ext != "" {
			_ = os.Remove("./avatar/" + s.UserName + s.Ext)
		}
		_ = c.SaveUploadedFile(avatar, "./avatar/"+s.UserName+filepath.Ext(avatar.Filename))
		s.Ext = filepath.Ext(avatar.Filename)
	}
	if newPassword != "" {
		s.Password = util.Md5Encode(newPassword)
	}
	if nickName != "" {
		s.NickName = nickName
	}
	model.UpdateSimpleUser(&s)
	c.HTML(http.StatusOK, "modify.html", gin.H{
		"status":   "success",
		"hint":     "修改个人信息成功",
		"userName": userName,
	})
}

// Login 登录
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"hint": "欢迎登录GO网盘！",
	})
}

// HandleSimpleLogin 处理简单用户登录
func HandleSimpleLogin(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")

	s := model.FindSimpleUserByUserName(userName)
	if s.UserName == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"hint": "用户名不存在！",
		})
		return
	}
	if util.Md5Encode(password) != s.Password {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"hint": "密码错误！",
		})
		return
	}
	// 生成token并写入cookie中
	token, err := util.GenerateToken(userName)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"hint": "无法生成token！",
		})
		return
	}
	c.SetCookie(
		"token",     // Cookie 名称
		token,       // Cookie 值
		3600*24,     // Cookie 的有效时间（秒）
		"/",         // Cookie 作用路径
		"localhost", // Cookie 作用域
		false,       // 是否开启 HTTPS
		true,        // 是否允许 JS 访问
	)
	// 获取当前登录时间并存储
	s.LastLoginTime = time.Now()
	model.UpdateSimpleUser(&s)
	//登录成功重定向到首页
	c.Redirect(http.StatusMovedPermanently, "/cloud/index")
}

func SimpleLogout(c *gin.Context) {
	// 通过设置过期时间为过去的时间清除Cookie
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"hint": "成功登出！",
	})
}
