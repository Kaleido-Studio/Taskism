package handlers

import (
	"net/http"
	"strings"

	"taskism/models"
	"taskism/pkgs/msg"
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
		msg.BadRequestErr(c)
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
		msg.BadRequestErr(c)
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		msg.InternalServerErr(c)
		return
	}
	user := models.NewUser(body.Name, string(password))
	if err := mgm.Coll(user).Create(user); err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			msg.ConflictErr(c)
			return
		}
		msg.BadRequestErr(c)
		return
	}
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		msg.BadRequestErr(c)
		return
	}
	c.JSON(http.StatusOK, RegisterResBody{Username: user.Name, Token: token})
}

type LoginResBodyJson struct {
	Token string `json:"token"`
}

func UserLoginHandler(c *gin.Context) {
	var body UserLogRegReqBody
	if err := c.ShouldBindJSON(&body); err != nil {
		msg.BadRequestErr(c)
		return
	}
	var user models.User
	err := mgm.Coll(&models.User{}).First(bson.M{"name": body.Name}, &user)
	if err != nil {
		msg.BadRequestErr(c)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		msg.BadRequestErr(c)
		return
	}
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		msg.InternalServerErr(c)
		return
	}
	c.JSON(http.StatusOK, LoginResBodyJson{Token: token})
}
