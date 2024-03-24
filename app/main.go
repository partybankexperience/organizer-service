package main

import (
	"github.com/djfemz/rave/app/security/controllers"
	"github.com/djfemz/rave/app/security/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", controllers.NewAuthController().LoginHandler)
	router.Use(middlewares.AuthMiddleware())
	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
