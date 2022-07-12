package main

import (
	"net/http"

	"taskism/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
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
	r.Run("localhost:3001")
}
