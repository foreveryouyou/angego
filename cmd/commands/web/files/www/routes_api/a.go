package routes_api

import "github.com/gin-gonic/gin"

// 前台api: http://{host}/api
func Route(r *gin.RouterGroup) {
	r.GET("/", todo)
}
