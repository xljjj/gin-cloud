package main

import (
	"CloudDrive/mysql"
	"CloudDrive/redis"
	"CloudDrive/router"
	"CloudDrive/util"
	"log"
)

func main() {
	util.InitConfig()
	mysql.InitMySQL()
	redis.InitRedis()

	r := router.Router()

	r.LoadHTMLGlob("view/*")        //配置页面路径
	r.Static("/static", "./static") //配置静态资源目录

	if err := r.Run(":8080"); err != nil {
		log.Fatal("无法启动服务器")
	}
}
