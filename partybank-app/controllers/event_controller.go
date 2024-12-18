package controllers

import (
	"errors"
	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/services"
	"github.com/djfemz/organizer-service/partybank-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

type EventController struct {
	services.EventService
	*validator.Validate
}

func NewEventController(eventService services.EventService, objectValidator *validator.Validate) *EventController {
	validator.New()
	return &EventController{
		eventService,
		objectValidator,
	}
}

// CreateEvent godoc
// @Summary      Add Event
// @Description  Adds Event
// @Tags         Events
// @Accept       json
// @Param 		 tags body dtos.CreateEventRequest true "Event tags"
// @Produce      json
// @Success      201  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event  [post]
func (eventController *EventController) CreateEvent(ctx *gin.Context) {
	createEventRequest := &request.CreateEventRequest{}
	err := ctx.BindJSON(createEventRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	err = eventController.Struct(createEventRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	res, err := eventController.EventService.Create(createEventRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &response.PartybankBaseResponse[*response.EventResponse]{Data: res})
}

// EditEvent godoc
// @Summary      Edit Event
// @Description  Edits Event
// @Tags         Events
// @Accept       json
// @Param 		 tags body dtos.UpdateEventRequest true "Event tags"
// @Param        id   path   int  true  "id"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/{id} [put]
func (eventController *EventController) EditEvent(ctx *gin.Context) {
	updateEventRequest := &request.UpdateEventRequest{}
	eventId, err := extractIdParamFromRequest("id", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	err = ctx.BindJSON(updateEventRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	err = eventController.Struct(updateEventRequest)
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

// GetAllEventsForSeries godoc
// @Summary      Get all Events in series
// @Description   Get all Events in series
// @Tags         Events
// @Accept       json
// @Param        seriesId  query   int  true  "seriesId"
// @Param        page   query   int  true  "page"
// @Param        size   query   int  true  "size"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/series [get]
func (eventController *EventController) GetAllEventsForSeries(ctx *gin.Context) {
	organizerId, err := strconv.ParseUint(ctx.Query("seriesId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return
	}
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageNumber, err := utils.ConvertQueryStringToInt(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	pageSize, err := utils.ConvertQueryStringToInt(size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	events, err := eventController.EventService.GetAllEventsFor(organizerId, pageNumber, pageSize)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// GetEventById godoc
// @Summary      Get Event By id
// @Description   Get Event By id
// @Tags         Events
// @Accept       json
// @Param        id  path int  true  "id"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/{id} [get]
func (eventController *EventController) GetEventById(ctx *gin.Context) {
	id, err := extractIdParamFromRequest("id", ctx)
	if err != nil {
		handleError(ctx, err)
	}
	event, err := eventController.EventService.GetById(id)
	if err != nil {
		handleError(ctx, err)
	}
	ctx.JSON(http.StatusOK, event)
}

// DiscoverEvents godoc
// @Summary      Discover events on the system
// @Description  Discover events on the system
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        page   query   int  true  "page"
// @Param        size   query   int  true  "size"
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Router       /api/v1/event/discover [get]
func (eventController *EventController) DiscoverEvents(ctx *gin.Context) {
	log.Println("In discover events")
	page := ctx.Query("page")
	pageNumber, err := utils.ConvertQueryStringToInt(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	size := ctx.Query("size")
	pageSize, err := utils.ConvertQueryStringToInt(size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	events, err := eventController.EventService.DiscoverEvents(pageNumber, pageSize)
	ctx.JSON(http.StatusOK, events)
}

// GetEventByReference godoc
// @Summary      Get Event By reference
// @Description   Get Event By reference
// @Tags         Events
// @Accept       json
// @Param        reference  path string  true  "reference"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Router       /api/v1/event/reference/{reference} [get]
func (eventController *EventController) GetEventByReference(ctx *gin.Context) {
	reference := ctx.Param("reference")
	event, err := eventController.EventService.GetEventByReference(reference)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

// PublishEvent godoc
// @Summary      Publish Event
// @Description  Publish Event
// @Tags         Events
// @Accept       json
// @Param        id  path int  true  "id"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/publish/{id} [get]
func (eventController *EventController) PublishEvent(ctx *gin.Context) {
	id, err := extractIdParamFromRequest("id", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	event, err := eventController.EventService.PublishEvent(id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, event)
}

// GetAllEventsForOrganizer  godoc
// @Summary      Get all Events for organizer
// @Description   Get all Events for organizer
// @Tags         Events
// @Accept       json
// @Param        organizerId  query   int  true  "organizerId"
// @Param        page   query   int  true  "page"
// @Param        size   query   int  true  "size"
// @Produce      json
// @Success      200  {object}  dtos.EventResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/organizer [get]
func (eventController *EventController) GetAllEventsForOrganizer(ctx *gin.Context) {
	organizerId, err := strconv.ParseUint(ctx.Query("organizerId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return
	}
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageNumber, err := utils.ConvertQueryStringToInt(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	pageSize, err := utils.ConvertQueryStringToInt(size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	events, err := eventController.EventService.GetAllEventsForOrganizer(organizerId, pageNumber, pageSize)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// DeleteEvent godoc
// @Summary      Delete Event by id
// @Description  Delete Event by id
// @Tags         Events
// @Accept       json
// @Param        eventId  path int  true  "eventId"
// @Produce      json
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/delete/{eventId} [delete]
func (eventController *EventController) DeleteEvent(ctx *gin.Context) {
	id, err := extractIdParamFromRequest("eventId", ctx)
	if err != nil {
		log.Println("error: ", err.Error())
		handleError(ctx, err)
		return
	}
	log.Println("event id: ", id)
	deleteEventResponse, err := eventController.EventService.DeleteEventBy(id)
	if err != nil {
		log.Println("error: ", err.Error())
		handleError(ctx, err)
		return
	}
	log.Println("delete response: ", deleteEventResponse)
	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[string]{Data: deleteEventResponse})
}

func (eventController *EventController) UpdateEventHasTicketSales(ctx *gin.Context) {
	eventReference := ctx.Query("eventReference")
	updateEventResponse, err := eventController.EventService.UpdateEventHasTicketSales(eventReference)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[string]{Data: updateEventResponse})
}

// UnpublishEvent godoc
// @Summary      Unpublish Event
// @Description  Unpublish Event
// @Tags         Events
// @Accept       json
// @Param        eventId  path int  true  "eventId"
// @Param 		 tags body dtos.UnPublishEventRequest true "Event tags"
// @Produce      json
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/event/unpublish/{eventId} [put]
func (eventController *EventController) UnpublishEvent(ctx *gin.Context) {
	eventId, err := extractIdParamFromRequest("eventId", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	unpublishRequest := &request.UnPublishEventRequest{}
	if err := ctx.BindJSON(unpublishRequest); err != nil {
		handleError(ctx, err)
		return
	}
	event, err := eventController.EventService.UnPublishEvent(eventId, unpublishRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func extractIdParamFromRequest(paramName string, ctx *gin.Context) (uint64, error) {
	log.Println("param name: ", paramName, "val: ", ctx.Param(paramName))
	id, err := strconv.ParseUint(ctx.Param(paramName), 10, 64)
	if err != nil {
		return 0, errors.New("error extracting path variable from request")
	}
	return id, err
}

func handleError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest,
		&response.PartybankBaseResponse[error]{Data: err})
	return
}
