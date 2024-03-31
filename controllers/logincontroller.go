package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ank809/JWT-Authentication-Go/database"
	"github.com/ank809/JWT-Authentication-Go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var jwt_key = []byte("secrey_key")

func Loginuser(c *gin.Context) {
	var user models.User
	var foundUser models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)

	filter1 := bson.M{"username": user.Username}
	err = collection.FindOne(context.Background(), filter1).Decode(&foundUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	expiration_time := time.Now().Add(time.Minute * 5)
	expected_password := foundUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(expected_password), []byte(user.Password))
	if err != nil {
		fmt.Println(expected_password)
		fmt.Print(user.Password)
		c.JSON(http.StatusBadRequest, "Invalid Password")
		return

	}
	claims := &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration_time,
	})

	// Send token back to the client
	c.JSON(http.StatusOK, gin.H{"token": tokenString,
		"success": "User loggin Successfully"})

}
