package handlers

import (
	"net/http"

	"taskism/structs"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	var person structs.Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"uuid": person.ID})
}
