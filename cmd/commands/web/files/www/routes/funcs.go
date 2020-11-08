package routes

import (
	"github.com/gin-gonic/gin"
)

func todo(c *gin.Context) {
	c.JSON(200, gin.H{
		"route": c.Request.RequestURI,
	})
}
