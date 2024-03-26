package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TicketController struct {
}

func NewTicketController() *TicketController {
	return &TicketController{}
}

func (ticketController *TicketController) AddTicketToEvent(ctx *gin.Context) {
	ticketService := services.NewTicketService()
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
