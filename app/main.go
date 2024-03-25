package main

import (
	security "github.com/djfemz/rave/app/security/controllers"
	"github.com/djfemz/rave/app/security/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	middlewares.Routers(router)
	authController := security.NewAuthController()
	router.POST("/auth/login", authController.AuthHandler)
	router.GET("/auth/validate-otp", authController.ValidateOtp)
	router.Use(middlewares.AuthMiddleware())
	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
