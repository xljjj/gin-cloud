package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Login 登录
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
