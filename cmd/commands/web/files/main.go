package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"

	"{{.ModuleName}}/global"
	"{{.ModuleName}}/www"
)

func main() {
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
		gin.SetMode(gin.ReleaseMode)
	}
	// 设置不同运行模式的配置
	if gin.Mode() == gin.ReleaseMode {
		log.Info("开始运行...")
		log.SetLevel(log.WarnLevel)
		log.SetFormatter(&log.JSONFormatter{})
		// 设置gin_log
		logFileName := fmt.Sprintf("./runtime/gin_http_%s.log", time.Now().Format("20060102150405"))
		logFile, err := os.Create(logFileName)
		if err != nil {
			log.Fatal("gin_http_log文件创建失败", err)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}
