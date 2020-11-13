package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"{{.ModuleName}}/global"
	"{{.ModuleName}}/www"
)

func main() {
	if global.SConf.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// Default 使用 Logger 和 Recovery 中间件
	app := gin.Default()
	www.RegisterRoutes(app)

	// 启动服务
	host := "0.0.0.0"
	port := global.SConf.Port
	addr := host + ":" + port
	err := app.Run(addr)
	if err != nil {
		log.Fatal("启动失败:", err)
	}
}

func init() {
	// do something here to set environment depending on an environment variable
	// or command-line flag
	// 初始化全局配置
	global.InitConfig("./conf.yml")
	if global.SConf.IsProd {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}
