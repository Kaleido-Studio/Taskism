package controllers

import (
	"net/http"

	"taskism/handlers"

	"github.com/gin-gonic/gin"
)

func GinEngine() *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"hello": "Yu Hanwen",
		}
		c.JSON(http.StatusOK, data)
	})
	api := r.Group("/api")
	{
		api.GET("/:id", handlers.GetHandler)
	}
	return r
}
