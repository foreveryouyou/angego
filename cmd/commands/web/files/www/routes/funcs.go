package routes

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/libs"
)

//BaseController
type BaseController struct {
	libs.BaseController
}

func todo(c *gin.Context) {
	c.JSON(200, gin.H{
		"route": c.Request.RequestURI,
	})
}
