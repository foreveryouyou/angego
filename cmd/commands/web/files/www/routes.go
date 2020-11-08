package www

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/www/middlewares"
	"{{.ModuleName}}/www/routes"
	"{{.ModuleName}}/www/routes_api"
)

// 注册路由
func RegisterRoutes(app *gin.Engine) {
	// favicon.ico
	app.StaticFile("/favicon.ico", "./static/favicon.ico")
	app.Static("/static", "./static")
	app.Use(middlewares.MWCors)

	//[default]
	routes.Route(app.Group("/"))

	//[api]
	routes_api.Route(app.Group("/api"))

	//[404]
	app.NoRoute(err404)
}

// 自定义404
func err404(ctx *gin.Context) {

}
