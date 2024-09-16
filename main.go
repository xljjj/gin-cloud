package main

import (
	"CloudDrive/mysql"
	"CloudDrive/redis"
	"CloudDrive/router"
	"CloudDrive/util"
	"html/template"
	"log"
	"strings"
)

func main() {
	util.InitConfig()
	mysql.InitMySQL()
	redis.InitRedis()

	r := router.Router()

	r.SetFuncMap(template.FuncMap{
		"concat": func(parts ...string) string { return strings.Join(parts, "") },
	}) //新增字符串拼接函数
	r.LoadHTMLGlob("view/*")        //配置页面路径
	r.Static("/static", "./static") //配置静态资源目录
	r.Static("/avatar", "./avatar")
	r.Static("/file", "./file") //静态文件服务，提供对 file 文件夹的访问

	if err := r.Run(":8080"); err != nil {
		log.Fatal("无法启动服务器")
	}
}
