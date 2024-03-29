package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Username string             `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
