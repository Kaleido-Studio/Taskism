package main

import (
	"net/http"

	"taskism/structs"

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
	r.GET("/:id", func(c *gin.Context) {
		var person structs.Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"uuid": person.ID})
	})
	r.Run("localhost:3001")
}
