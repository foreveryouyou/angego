package middlewares

import "github.com/gin-contrib/cors"

var MWCors = cors.New(cors.Config{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
		"HEAD"},
	AllowHeaders: []string{
		"Origin",
		"Content-Type",
		"Cookie",
		"Accept",
		"Authorization",
		"Access-Control-Allow-Origin",
		"X-Requested-With",
		"X-Sys-Id",
	},
	AllowCredentials: true,
})
