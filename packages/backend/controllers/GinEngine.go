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
	apiUser := r.Group("/api/user")
	{
		apiUser.GET("/info/:id", handlers.UserGetHandler)
		apiUser.POST("/register", handlers.UserRegisterHandler)
		apiUser.POST("/login", handlers.UserLoginHandler)
	}
	apiProject := r.Group("/api/project")
	{
		apiProject.GET("/", handlers.ProjectGetHandler)
		apiProject.POST("/", handlers.ProjectPostHandler)
		apiProject.PUT("/:id", handlers.ProjectPutHandler)
		apiProject.DELETE("/:id", handlers.ProjectDeleteHandler)
	}
	return r
}
