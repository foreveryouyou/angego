package routes

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/global"
)

// 前台: http://{host}/
func Route(r *gin.RouterGroup) {
	// index
	r.GET("/", RouteA{}.Index)
	// version
	r.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"buildTime":      global.BuildTime,
			"buildGoVersion": global.BuildGoVersion,
		})
	})
}
