package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventController struct {
	services.EventService
}

func NewEventController() *EventController {
	return &EventController{
		services.NewEventService(),
	}
}

func (eventController *EventController) EditEvent(ctx *gin.Context) {
	updateEventRequest := &request.UpdateEventRequest{}
	eventId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			&response.RaveResponse[string]{Data: err.Error()})
		return
	}
	err = ctx.BindJSON(updateEventRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			&response.RaveResponse[string]{Data: err.Error()})
		return
	}
	updateEventResponse, err := eventController.EventService.UpdateEventInformation(eventId, updateEventRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			&response.RaveResponse[string]{Data: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updateEventResponse)
}
