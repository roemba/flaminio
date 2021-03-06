package main

import (
	"flaminio/handlers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./public", true))) //May have some nasty performance implications
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", handlers.LoginHandler)
		v1.GET("/locations/*uuid", handlers.GETLocationsHandler)
		authorized := v1.Group("")
		authorized.Use(validateTokenMiddleware)
		{
			authorized.GET("/auth/user", handlers.UserHandler)
			authorized.GET("/auth/refresh", handlers.RefreshHandler)
			authorized.POST("/auth/logout", handlers.LogoutHandler)

			authorized.GET("/reservations", handlers.GETReservationsHandler)

			authorized.DELETE("/locations/:uuid", handlers.DELETELocationHandler)

			jsonRequestBody := authorized.Group("")
			jsonRequestBody.Use(checkMediaTypeHeaderIsJson)
			{
				jsonRequestBody.POST("/reservations", handlers.POSTReservationsHandler)
				jsonRequestBody.POST("/locations", handlers.POSTLocationsHandler)
				jsonRequestBody.PUT("/locations", handlers.PUTLocationsHandler)
			}
		}
	}
}
