package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonGet struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func GetHandler(c *gin.Context) {
	var person PersonGet
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uuid": person.ID})
}
