package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TicketController struct {
}

var ticketService = services.NewTicketService()

func NewTicketController() *TicketController {
	return &TicketController{}
}

// AddTicketToEvent godoc
// @Summary      Add Ticket to Event
// @Description  Add Ticket to Event
// @Tags         Tickets
// @Accept       json
// @Param 		 tags body dtos.CreateTicketRequest true "Ticket tags"
// @Produce      json
// @Success      200  {object}  dtos.TicketResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /protected/ticket [post]
func (ticketController *TicketController) AddTicketToEvent(ctx *gin.Context) {
	addTicketRequest := &request.CreateTicketRequest{}
	err := ctx.BindJSON(addTicketRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response, err := ticketService.CreateTicketFor(addTicketRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// GetAllTicketsForEvent godoc
// @Summary      Get all Tickets for Event
// @Description   Get all Tickets for Event
// @Tags         Tickets
// @Accept       json
// @Param        eventId path int  true  "eventId"
// @Produce      json
// @Success      200  {array}  models.Ticket
// @Failure      400  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /protected/ticket/{eventId} [get]
func (ticketController *TicketController) GetAllTicketsForEvent(ctx *gin.Context) {
	eventId, err := extractParamFromRequest("eventId", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	tickets, err := ticketService.GetAllTicketsFor(eventId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

// GetTicketById godoc
// @Summary      Get Ticket By id
// @Description   Get Ticket By id
// @Tags         Tickets
// @Accept       json
// @Param        ticketId  query int  true  "ticketId"
// @Produce      json
// @Success      200  {object}  dtos.TicketResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /protected/ticket [get]
func (ticketController *TicketController) GetTicketById(ctx *gin.Context) {
	eventId, err := strconv.ParseUint(ctx.Query("ticketId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ticket, err := ticketService.GetTicketById(eventId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, ticket)
}
