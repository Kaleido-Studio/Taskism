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
	api := r.Group("/api")
	userRoute := api.Group("/user")
	{
		userRoute.GET("/info/:id", handlers.UserGetHandler)
		userRoute.POST("/register", handlers.UserRegisterHandler)
		userRoute.POST("/login", handlers.UserLoginHandler)
	}
	projectRouter := api.Group("/project")
	{
		projectRouter.GET("/", handlers.ProjectGetHandler)
		projectRouter.POST("/", handlers.ProjectPostHandler)
		projectRouter.PUT("/:id", handlers.ProjectPutHandler)
		projectRouter.DELETE("/:id", handlers.ProjectDeleteHandler)
	}
	return r
}
