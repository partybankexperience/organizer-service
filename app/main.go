package main

import (
	"github.com/djfemz/rave/app/security"
	"github.com/spf13/viper"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	viper.BindEnv("")
	router := gin.Default()

	router.GET("", security.LoginHandler)

	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
