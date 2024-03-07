package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	router := gin.Default()

	router.GET("")

	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}






