package middleware

import (
	"CloudDrive/model"
	"CloudDrive/redis"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// CheckLogin 检查登录中间件
func CheckLogin(c *gin.Context) {
	token, err := c.Cookie("Token")
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}
	openId, err := redis.GetKey(c, token)
	if err != nil {
		fmt.Println("Get Redis Err:", err.Error())
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	user := model.GetUser(openId)

	if user.Id == 0 {
		//校验失败 返回登录页面
		c.Redirect(http.StatusFound, "/")
	} else {
		//校验成功 继续执行
		c.Set("openId", openId)
		c.Next()
	}
}

func CheckSimpleLogin(c *gin.Context) {
	// 从请求的 Cookie 中获取 Token
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hint": "请先登录！"})
		c.Abort()
		return
	}

	claims, err := util.ParseToken(tokenString)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hint": "请重新登录！"})
		c.Abort()
		return
	}

	// 将用户名及用户ID存储在上下文中
	user := model.FindSimpleUserByUserName(claims.UserName)
	c.Set("userName", user.UserName)
	c.Set("userId", user.Id)
	c.Next()
}

func CheckAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hint": "请先登录！"})
		c.Abort()
		return
	}

	claims, err := util.ParseToken(tokenString)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hint": "请重新登录！"})
		c.Abort()
		return
	}

	if claims.UserName != "admin" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hint": "请使用管理员账号登录！"})
		c.Abort()
		return
	}

	// 将用户名及用户ID存储在上下文中
	user := model.FindSimpleUserByUserName(claims.UserName)
	c.Set("userName", user.UserName)
	c.Next()
}

// CheckInput 拦截用户直接访问文件
func CheckInput(c *gin.Context) {
	curPath := c.Request.URL.Path
	if len(curPath) >= 6 && curPath[:6] == "/file/" {
		// 用户登录后才能访问文件
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{"hint": "请先登录！"})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{"hint": "请先登录！"})
			c.Abort()
			return
		}
		// 检查是否是该用户的仓库
		trimmedPath := strings.TrimPrefix(curPath, "/file/")
		parts := strings.SplitN(trimmedPath, "/", 2)
		if len(parts) < 2 {
			// 如果分割后没有两部分，说明路径无效
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"hint": "无效路径，请登录后在页面中查看！"})
			c.Abort()
			return
		}
		storeId := parts[0]
		user := model.FindSimpleUserByUserName(claims.UserName)
		if strconv.Itoa(user.FileStoreId) != storeId {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"hint": "禁止访问他人的文件！"})
			c.Abort()
			return
		}
	}
	c.Next()
}
