package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type OrganizerController struct {
	organizerService services.OrganizerService
	*validator.Validate
}

func NewOrganizerController(organizerService services.OrganizerService, objectValidator *validator.Validate) *OrganizerController {
	return &OrganizerController{
		organizerService,
		objectValidator,
	}
}

// AddEventStaff godoc
// @Summary      Add Event Staff
// @Description  Adds Event Staff
// @Tags         Organizer
// @Accept       json
// @Param 		 tags body dtos.AddEventStaffRequest true "Organizer tags"
// @Produce      json
// @Success      200  {object}  dtos.RaveResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /api/v1/event/staff [post]
func (orgController *OrganizerController) AddEventStaff(ctx *gin.Context) {
	addEventStaff := &request.AddEventStaffRequest{}
	err := ctx.BindJSON(addEventStaff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
		return
	}
	err = orgController.Struct(addEventStaff)
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
