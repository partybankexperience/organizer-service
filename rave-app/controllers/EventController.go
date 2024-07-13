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

func (eventController *EventController) CreateEvent(ctx *gin.Context) {
	createEventRequest := &request.CreateEventRequest{}
	err := ctx.BindJSON(createEventRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
		return
	}
	res, err := eventController.EventService.Create(createEventRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (eventController *EventController) EditEvent(ctx *gin.Context) {
	updateEventRequest := &request.UpdateEventRequest{}
	eventId, err := extractParamFromRequest("id", ctx)
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

func (eventController *EventController) GetAllEventsForOrganizer(ctx *gin.Context) {
	organizerId, err := strconv.ParseUint(ctx.Query("organizerId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return
	}
	events, err := eventController.EventService.GetAllEventsFor(organizerId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func handleError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest,
		&response.RaveResponse[string]{Data: err.Error()})
	return
}

func extractParamFromRequest(paramName string, ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Param(paramName), 10, 64)
}

func (eventController *EventController) GetEventById(ctx *gin.Context) {
	id, err := extractParamFromRequest("id", ctx)
	if err != nil {
		handleError(ctx, err)
	}
	event, err := eventController.EventService.GetById(id)
	if err != nil {
		handleError(ctx, err)
	}
	ctx.JSON(http.StatusOK, event)
}
