package controllers

import (
	"log"
	"mime"
	"os"

	"taskism/handlers"
	"taskism/middlewares"

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
	r.Use(middlewares.Spa())
	mime.AddExtensionType(".js", "application/javascript")
	api := r.Group("/api/user")
	{
		api.GET("/info/:id", handlers.UserGetHandler)
		api.POST("/register", handlers.UserRegisterHandler)
		api.POST("/login", handlers.UserLoginHandler)
	}
	return r
}
