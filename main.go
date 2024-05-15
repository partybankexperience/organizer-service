package main

import (
	security "github.com/djfemz/rave/rave-app/security/controllers"
	"github.com/djfemz/rave/rave-app/security/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error starting server: ", err)
	}
	router := gin.Default()
	middlewares.Routers(router)
	router.Use(cors.New(configureCors()))
	authController := security.NewAuthController()
	router.POST("/auth/login", authController.AuthHandler)
	router.GET("/auth/validate-otp", authController.ValidateOtp)
	err = router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}

func configureCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{http.MethodOptions,
		http.MethodPost, http.MethodOptions, http.MethodPost, http.MethodGet}
	return config
}
