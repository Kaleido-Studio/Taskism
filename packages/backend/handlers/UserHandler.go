package handlers

import (
	"log"
	"net/http"
	"strings"

	"taskism/models"
	"taskism/utils"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
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

type UserLogRegReqBody struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResBody struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func UserRegisterHandler(c *gin.Context) {
	var body UserLogRegReqBody
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
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, RegisterResBody{Username: user.Name, Token: token})
}

type LoginResBodyJson struct {
	Token string `json:"token"`
}

func UserLoginHandler(c *gin.Context) {
	var body UserLogRegReqBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Panicln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	err := mgm.Coll(&models.User{}).First(bson.M{"name": body.Name}, &user)
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		log.Panicln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, LoginResBodyJson{Token: token})
}
