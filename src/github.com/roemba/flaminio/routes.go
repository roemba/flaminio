package flaminio

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func setRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./public", true))) //May have some nasty performance implications
	router.NoRoute(static.Serve("/", static.LocalFile("./public", true)))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", LoginHandler)
		authorized := v1.Group("/")
		authorized.Use(ValidateTokenMiddleware)
		{
			authorized.GET("/auth/user", UserHandler)
			authorized.GET("/auth/refresh", RefreshHandler)
		}
	}
}
