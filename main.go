package main

import (
	"github.com/ank809/JWT-Authentication-Go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", helloworld)
	router.GET("/signup", controllers.SignupUser)
	router.GET("/login", controllers.Loginuser)
	router.GET("/home", controllers.HomePage)
	router.Run(":8081")
}

func helloworld(c *gin.Context) {
	c.JSON(200, gin.H{
		"Succcess": "HELLO WORLD!",
	})
}
