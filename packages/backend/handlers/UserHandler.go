package handlers

import (
	"net/http"
	"strings"

	"taskism/models"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type PersonGet struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func UserGetHandler(c *gin.Context) {
	var person PersonGet
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uuid": person.ID})
}

type UserRegisterPostBody struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRegisterHandler(c *gin.Context) {
	var body UserRegisterPostBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.NewUser(body.Name, string(password))
	if err := mgm.Coll(user).Create(user); err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
