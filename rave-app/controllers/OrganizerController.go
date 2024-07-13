package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizerController struct {
	organizerService services.OrganizerService
}

func NewOrganizerController() *OrganizerController {
	return &OrganizerController{
		services.NewOrganizerService(),
	}
}

func (orgController *OrganizerController) AddEventStaff(ctx *gin.Context) {
	addEventStaff := &request.AddEventStaffRequest{}
	err := ctx.BindJSON(addEventStaff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
		return
	}
	res, err := orgController.organizerService.AddEventStaff(addEventStaff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
