package middlewares

import (
	control "github.com/djfemz/rave/rave-app/controllers"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/security"
	"github.com/djfemz/rave/rave-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Routers(router *gin.Engine) {
	organizerController := control.NewOrganizerController()
	protected := router.Group("/protected")
	{
		protected.POST("/event", organizerController.CreateEvent)
		protected.POST("/event-staff", organizerController.AddEventStaff)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
		authValue := strings.Split(" ", authHeader)
		token := authValue[len(authValue)-1]
		org, err := security.ExtractUserFrom(token)
		if err != nil {
			ctx.JSON(http.StatusForbidden, &response.RaveResponse[string]{Data: "token is invalid"})
			return
		}
		if org != nil {
			ctx.Next()
		}
	}
}
