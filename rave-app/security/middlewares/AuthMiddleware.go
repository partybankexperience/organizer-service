package middlewares

import (
	handlers "github.com/djfemz/rave/rave-app/controllers"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/security"
	"github.com/djfemz/rave/rave-app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var routesAuthorities map[string][]string

func Routers(router *gin.Engine) {
	organizerController := handlers.NewOrganizerController()
	eventController := handlers.NewEventController()
	ticketController := handlers.NewTicketController()
	calendarController := handlers.NewCalendarController()
	protected := router.Group("/protected", AuthMiddleware())
	{
		protected.POST("/event", organizerController.CreateEvent)
		protected.GET("/event/:id", eventController.GetEventById)
		protected.PUT("/event/:id", eventController.EditEvent)
		protected.GET("/event", eventController.EditEvent)
		protected.GET("/event/organizer", eventController.GetAllEventsForOrganizer)
		protected.POST("/event/staff", organizerController.AddEventStaff)
		protected.POST("/ticket", ticketController.AddTicketToEvent)
		protected.GET("/ticket/:eventId", ticketController.GetAllTicketsForEvent)
		protected.GET("/ticket", ticketController.GetTicketById)
		protected.POST("/calendar", calendarController.CreateCalendar)
		protected.GET("/calendar/:id", calendarController.GetCalendar)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
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
	return token
}
