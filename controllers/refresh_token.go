package controllers

import (
	"net/http"
	"time"

	"github.com/ank809/JWT-Authentication-Go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	tokenstr := cookie

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
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
		c.JSON(http.StatusBadRequest, "Token invalid")
	}
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30 {
	// c.JSON(http.StatusBadRequest, "Token is valid ")
	// return
	// }

	expiration_time := time.Now().Add(time.Minute * 5)
	// if time.Until(expiration_time) > time.Second*30 {
	// 	c.JSON(http.StatusBadRequest, "Token is valid ")
	// 	return
	// }
	claims.ExpiresAt = expiration_time.Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := tkn.SignedString(jwt_key)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenstring,
		Expires: expiration_time,
	})

	c.JSON(200, gin.H{
		"refresh_token": tokenstring,
		"message":       "success",
	})
}
