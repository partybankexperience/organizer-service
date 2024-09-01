package main

import (
	_ "github.com/djfemz/rave/docs"
	handlers "github.com/djfemz/rave/rave-app/controllers"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/security/middlewares"
	services2 "github.com/djfemz/rave/rave-app/security/services"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

var err error
var db *gorm.DB
var eventRepository repositories.EventRepository
var organizerRepository repositories.OrganizerRepository
var seriesRepository repositories.SeriesRepository
var ticketRepository repositories.TicketRepository
var eventStaffRepository repositories.EventStaffRepository

var eventService services.EventService
var organizerService services.OrganizerService
var seriesService services.SeriesService
var ticketService services.TicketService
var eventStaffService services.EventStaffService
var authService *services2.AuthService

var organizerController *handlers.OrganizerController
var eventController *handlers.EventController
var seriesController *handlers.SeriesController
var ticketController *handlers.TicketController

var objectValidator *validator.Validate

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading configuration: ", err)
	}
	log.Println("connnecting to db")
	db = repositories.Connect()
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
// @host 		  rave.onrender.com
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
func main() {
	router := gin.Default()
	configureAppComponents()
	middlewares.Routers(router, organizerController,
		eventController, seriesController, ticketController,
		authService)

	err = router.Run(":8000")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}

func configureAppComponents() {
	eventRepository = repositories.NewEventRepository(db)
	organizerRepository = repositories.NewOrganizerRepository(db)
	seriesRepository = repositories.NewSeriesRepository(db)
	ticketRepository = repositories.NewTicketRepository(db)
	eventStaffRepository = repositories.NewEventStaffRepository(db)

	seriesService = services.NewSeriesService(seriesRepository)
	eventStaffService = services.NewEventStaffService(eventStaffRepository, eventRepository)
	organizerService = services.NewOrganizerService(organizerRepository, eventStaffService, seriesService)
	eventService = services.NewEventService(eventRepository, organizerService, seriesService)
	ticketService = services.NewTicketService(ticketRepository, eventService)
	authService = services2.NewAuthService(organizerService, services.NewMailService())
	objectValidator = validator.New()

	organizerController = handlers.NewOrganizerController(organizerService, objectValidator)
	eventController = handlers.NewEventController(eventService, objectValidator)
	seriesController = handlers.NewSeriesController(seriesService, objectValidator)
	ticketController = handlers.NewTicketController(ticketService, objectValidator)
}
