package main

import (
	_ "github.com/djfemz/organizer-service/docs"
	"github.com/djfemz/organizer-service/partybank-app/initiator"
)

//organizer-service.onrender.com

// @title           Partybank Organizer Service
// @version         1.0
// @description     Partybank Organizer Service.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    https://www.thepartybank.com
// @contact.email  unavailable
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host organizer-service.onrender.com
// @schemes https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
func main() {
	//organizer-service.onrender.com
	initiator.Initiate()
}
