package routes

import (
	"github.com/gin-gonic/gin"
)

type RouteA struct {
	BaseController
}

func (p RouteA) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}
