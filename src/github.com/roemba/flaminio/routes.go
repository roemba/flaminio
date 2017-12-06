package flaminio

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/roemba/flaminio/handlers"
)

func setRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./public", true))) //May have some nasty performance implications
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", handlers.LoginHandler)
		authorized := v1.Group("/")
		authorized.Use(validateTokenMiddleware)
		{
			authorized.GET("/auth/user", handlers.UserHandler)
			authorized.GET("/auth/refresh", handlers.RefreshHandler)
			authorized.GET("/reservations", handlers.GETReservationsHandler)
			authorized.PUT("/reservations", handlers.PUTReservationsHandler)
			authorized.PUT("/locations", handlers.PUTLocationsHandler)
		}
	}
}
