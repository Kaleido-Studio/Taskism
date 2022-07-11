package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"hello": "Yu Hanwen",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run("localhost:3001")
}
