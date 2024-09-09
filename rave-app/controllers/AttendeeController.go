package controllers

import (
	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AttendeeController struct {
	attendeeService services.AttendeeService
	objectValidator *validator.Validate
}

func NewAttendeeController(attendeeService services.AttendeeService, objectValidator *validator.Validate) *AttendeeController {
	return &AttendeeController{
		attendeeService: attendeeService,
		objectValidator: objectValidator,
	}
}

// Register godoc
// @Summary      Onboard Attendee
// @Description  Onboard Attendee
// @Tags         Auth
// @Accept       json
// @Param 		 tags body dtos.CreateAttendeeRequest true "Auth tags"
// @Produce      json
// @Success      201  {object}  dtos.RaveResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Router       /auth/attendee  [post]
func (attendeeController AttendeeController) Register(ctx *gin.Context) {
	attendeeAuthRequest := &dtos.CreateAttendeeRequest{}
	err := ctx.BindJSON(attendeeAuthRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	err = attendeeController.objectValidator.Struct(attendeeAuthRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	res, err := attendeeController.attendeeService.Register(attendeeAuthRequest)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.RaveResponse[*response.AttendeeResponse]{Data: res})
}
