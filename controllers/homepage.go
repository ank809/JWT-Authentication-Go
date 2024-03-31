package controllers

import (
	"net/http"

	"github.com/ank809/JWT-Authentication-Go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	tokenStr := cookie

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, err
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if !token.Valid {
		c.JSON(http.StatusBadRequest, "Token is invalid")
		return
	}
	c.JSON(200, gin.H{
		"success": "you are authorized",
		"token":   tokenStr,
	})
}
