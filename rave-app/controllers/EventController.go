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
	eventId, err := extractIdFrom("id", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	err = ctx.BindJSON(updateEventRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	updateEventResponse, err := eventController.EventService.UpdateEventInformation(eventId, updateEventRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, updateEventResponse)
}

func handleError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest,
		&response.RaveResponse[string]{Data: err.Error()})
	return
}

func extractIdFrom(key string, ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Param(key), 10, 64)
}

func (eventController *EventController) GetEventById(ctx *gin.Context) {
	id, err := extractIdFrom("id", ctx)
	if err != nil {
		handleError(ctx, err)
	}
	event, err := eventController.EventService.GetById(id)
	if err != nil {
		handleError(ctx, err)
	}
	ctx.JSON(http.StatusOK, event)
}
