package handlers

import (
	"net/http"

	"taskism/models"
	"taskism/utils/msg"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

type NewProjectReq struct {
	Name string `json:"name" binding:"required"`
}

func ProjectGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func ProjectPostHandler(c *gin.Context) {
	var newProj NewProjectReq
	if err := c.ShouldBindJSON(&newProj); err != nil {
		msg.BadRequestErr(c)
	}
	project := models.NewProject(newProj.Name)
	if err := mgm.Coll(project).Create(project); err != nil {
		msg.BadRequestErr(c)
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func ProjectPutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func ProjectDeleteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
