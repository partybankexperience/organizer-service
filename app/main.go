package main

import (
	"github.com/djfemz/rave/app/security"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", security.LoginHandler)

	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
