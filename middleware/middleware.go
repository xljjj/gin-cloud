package middleware

import (
	"CloudDrive/model"
	"CloudDrive/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
