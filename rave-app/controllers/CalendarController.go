package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalendarController struct {
	services.CalendarService
}

func NewCalendarController() *CalendarController {
	return &CalendarController{services.NewCalendarService()}
}

func (calendarController *CalendarController) CreateCalendar(ctx *gin.Context) {
	createCalendarRequest := &request.CreateCalendarRequest{}
	calendarService := calendarController.CalendarService
	err := ctx.BindJSON(createCalendarRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := calendarService.CreateCalendar(createCalendarRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, &response.RaveResponse[*response.CreateCalendarResponse]{resp})

}

func (calendarController *CalendarController) GetCalendar(ctx *gin.Context) {
	calendarService := calendarController.CalendarService
	id, err := extractParamFromRequest("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	calendar, err := calendarService.GetCalendar(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, &response.RaveResponse[*response.CreateCalendarResponse]{calendar})

}
