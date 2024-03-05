package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("")

	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
