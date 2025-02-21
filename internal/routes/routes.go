package routes

import (
	"booking/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", handlers.Home)

	return router
}
