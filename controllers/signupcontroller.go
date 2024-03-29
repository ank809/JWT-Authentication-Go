package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ank809/JWT-Authentication-Go/database"
	"github.com/ank809/JWT-Authentication-Go/helpers"
	"github.com/ank809/JWT-Authentication-Go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(c *gin.Context) {
	var user models.User
	user.ID = primitive.NewObjectID()
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("Error in binding json", err)
		return
	}

	var hashed_password []byte
	isPasswordValid, value := helpers.CheckPassword(user.Password)
	if !isPasswordValid {
		c.JSON(http.StatusBadRequest, value)
		return
	}
	hashed_password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed_password)
	isUsernameValid, str := helpers.CheckUsername(user.Username)
	if !isUsernameValid {
		fmt.Println(str)
		return
	}

	isEmailValid, msg := helpers.VerifyEMail(user.Email)
	if !isEmailValid {
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	collection.InsertOne(context.Background(), user)
	c.JSON(200, "User added successfully")

}
