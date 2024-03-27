package main

import (
	security "github.com/djfemz/rave/rave-app/security/controllers"
	"github.com/djfemz/rave/rave-app/security/middlewares"
	"github.com/joho/godotenv"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	router := gin.Default()
	middlewares.Routers(router)
	authController := security.NewAuthController()
	router.POST("/auth/login", authController.AuthHandler)
	router.GET("/auth/validate-otp", authController.ValidateOtp)
	err = router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
