package controllers

import (
	"log"
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
	api := r.Group("/api/user")
	{
		api.GET("/info/:id", handlers.UserGetHandler)
		api.POST("/register", handlers.UserRegisterHandler)
	}
	return r
}
