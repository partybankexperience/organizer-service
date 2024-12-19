package initiator

import (
	handlers "github.com/djfemz/organizer-service/partybank-app/controllers"
	"github.com/djfemz/organizer-service/partybank-app/integrations"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
	"github.com/djfemz/organizer-service/partybank-app/security/controllers"
	services2 "github.com/djfemz/organizer-service/partybank-app/security/services"
	"github.com/djfemz/organizer-service/partybank-app/services"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"os"
)

var err error
var db *gorm.DB
var eventRepository repositories.EventRepository
var organizerRepository repositories.OrganizerRepository
var seriesRepository repositories.SeriesRepository
var ticketRepository repositories.TicketRepository
var eventStaffRepository repositories.EventStaffRepository
var attendeeRepository repositories.AttendeeRepository
var imageRepository repositories.ImageRepository

var eventService services.EventService
var organizerService services.OrganizerService
var seriesService services.SeriesService
var ticketService services.TicketService
var eventStaffService services.EventStaffService
var authService *services2.AuthService
var attendeeService services.AttendeeService
var imageService services.ImageService
var fileService services.FileUploadService

var organizerController *handlers.OrganizerController
var eventController *handlers.EventController
var seriesController *handlers.SeriesController
var ticketController *handlers.TicketController
var attendeeController *handlers.AttendeeController
var authController *controllers.AuthController
var paymentService integrations.PaymentService
var objectValidator *validator.Validate
var imageController *handlers.FileController

func configureAppComponents() {
	objectValidator = validator.New()
	configureRepositoryComponents()
	configureServiceComponents()
	configureControllers()
}

func configureControllers() {
	organizerController = handlers.NewOrganizerController(organizerService, objectValidator)
	eventController = handlers.NewEventController(eventService, objectValidator)
	seriesController = handlers.NewSeriesController(seriesService, objectValidator)
	ticketController = handlers.NewTicketController(ticketService, objectValidator)
	attendeeController = handlers.NewAttendeeController(attendeeService, objectValidator)
	authController = controllers.NewAuthController(authService)
	imageController = handlers.NewFileController(fileService)
}

func configureServiceComponents() {
	mailService := services.NewGoMailService()
	seriesService = services.NewSeriesService(seriesRepository)
	eventStaffService = services.NewEventStaffService(eventStaffRepository, eventRepository)
	organizerService = services.NewOrganizerService(organizerRepository, eventStaffService, seriesService, ticketService, attendeeService)
	paymentService = integrations.NewPaymentService(os.Getenv("PAYMENT_SERVICE_CLIENT_ID"), os.Getenv("PAYMENT_SERVICE_CLIENT_SECRET"))
	eventService = services.NewEventService(eventRepository, organizerService, seriesService, ticketService, paymentService)
	seriesService.SetEventService(eventService)
	ticketService = services.NewTicketService(ticketRepository, eventService)
	eventService.SetTicketService(ticketService)
	attendeeService = services.NewAttendeeService(attendeeRepository, mailService)
	authService = services2.NewAuthService(organizerService, attendeeService, mailService)
	imageService = services.NewImageService(imageRepository)
	fileService = services.NewFileUploadService(imageService)
}

func configureRepositoryComponents() {
	eventRepository = repositories.NewEventRepository(db)
	organizerRepository = repositories.NewOrganizerRepository(db)
	seriesRepository = repositories.NewSeriesRepository(db)
	ticketRepository = repositories.NewTicketRepository(db)
	eventStaffRepository = repositories.NewEventStaffRepository(db)
	attendeeRepository = repositories.NewAttendeeRepository(db)
	imageRepository = repositories.NewPartybankImageRepository(db)
}
