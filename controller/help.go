package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Help(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", gin.H{})
}
