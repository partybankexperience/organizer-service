package main

import (
	_ "github.com/djfemz/rave/docs"
	"github.com/djfemz/rave/rave-app/security/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

var err error

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading configuration: ", err)
	}
}

// @title           Partybank Organizer Service
// @version         1.0
// @description     Partybank Organizer Service.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    https://www.thepartybank.com
// @contact.email  unavailable
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8082
// @BasePath  /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	middlewares.Routers(router)

	err = router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
