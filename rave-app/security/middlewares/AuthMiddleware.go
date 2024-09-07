package middlewares

import (
	handlers "github.com/djfemz/rave/rave-app/controllers"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/security"
	"github.com/djfemz/rave/rave-app/security/controllers"
	"github.com/djfemz/rave/rave-app/security/services"
	"github.com/djfemz/rave/rave-app/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strings"
)

var routesAuthorities map[string][]string

func Routers(router *gin.Engine, organizerController *handlers.OrganizerController,
	eventController *handlers.EventController, seriesController *handlers.SeriesController,
	ticketController *handlers.TicketController, authService *services.AuthService) {

	protected := router.Group("/api/v1", AuthMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		protected.POST("/event", eventController.CreateEvent)
		protected.GET("/event/:id", eventController.GetEventById)
		protected.PUT("/event/:id", eventController.EditEvent)
		protected.GET("/event/organizer", eventController.GetAllEventsForOrganizer)
		protected.POST("/event/staff", organizerController.AddEventStaff)
		protected.POST("/ticket", ticketController.AddTicketToEvent)
		protected.GET("/ticket/:eventId", ticketController.GetAllTicketsForEvent)
		protected.GET("/ticket", ticketController.GetTicketById)
		protected.POST("/series", seriesController.CreateSeries)
		protected.GET("/series/:id", seriesController.GetSeriesById)
		protected.GET("/series/organizer/:organizerId", seriesController.GetSeriesForOrganizer)
		protected.GET("/api/v1/event/publish/:id", eventController.PublishEvent)
	}
	router.Use(cors.New(configureCors()))
	authController := controllers.NewAuthController(authService)
	router.POST("/auth/login", authController.AuthHandler)
	router.GET("/auth/otp/validate", authController.ValidateOtp)
	router.GET("/api/v1/event/discover", eventController.DiscoverEvents)
	router.GET("/api/v1/event/reference/:reference", eventController.GetEventByReference)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
		log.Println("auth header: ", authHeader)
		token := extractTokenFrom(authHeader)
		if !isValid(token) {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				&response.RaveResponse[string]{Data: "access token is invalid"})
			return
		}
		ctx.Next()
	}
}

func isValid(token string) bool {
	org, err := security.ExtractUserFrom(token)
	log.Println("\norg: ", org, "\nerr: ", err)
	if err != nil || org == nil {
		return false
	}
	return true
}

func extractTokenFrom(authHeader string) string {
	authValue := strings.Split(authHeader, " ")
	token := authValue[len(authValue)-1]
	log.Println("token: ", token)
	return token
}

func configureCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{http.MethodOptions,
		http.MethodPost, http.MethodOptions, http.MethodPost, http.MethodGet}
	return config
}
