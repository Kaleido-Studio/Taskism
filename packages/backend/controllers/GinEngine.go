package controllers

import (
	"log"
	"net/http"
	"os"

	"taskism/handlers"

	"github.com/gin-gonic/gin"
)

func GinEngine() *gin.Engine {
	InitDB()
	log.SetPrefix("[Taskism] ")
	if os.Getenv("CI") != "" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"hello": "Yu Hanwen",
		}
		c.JSON(http.StatusOK, data)
	})
	api := r.Group("/api/user")
	{
		api.GET("/find/:id", handlers.UserGetHandler)
		api.POST("/register", handlers.UserRegisterHandler)
	}
	return r
}
