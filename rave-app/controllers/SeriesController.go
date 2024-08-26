package controllers

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SeriesController struct {
	services.SeriesService
	*validator.Validate
}

func NewSeriesController() *SeriesController {
	return &SeriesController{
		services.NewSeriesService(),
		validator.New(),
	}
}

// CreateSeries godoc
// @Summary      Add New Series
// @Description  Adds New Series
// @Tags         Series
// @Accept       json
// @Param 		 tags body dtos.CreateSeriesRequest true "Series tags"
// @Produce      json
// @Success      200  {object}  dtos.RaveResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /api/v1/series [post]
func (seriesController *SeriesController) CreateSeries(ctx *gin.Context) {
	createSeriesRequest := &request.CreateSeriesRequest{}
	err := ctx.BindJSON(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
		return
	}
	err = seriesController.Struct(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
		return
	}
	resp, err := seriesController.AddSeries(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
		return
	}
	ctx.JSON(http.StatusCreated, &response.RaveResponse[*response.CreateCalendarResponse]{Data: resp})
}

// GetSeriesById godoc
// @Summary      Get Series by id
// @Description  Get Series by id
// @Tags         Series
// @Accept       json
// @Produce      json
// @Param        id   path   int  true  "series id"
// @Success      200  {object}  dtos.RaveResponse
// @Failure      400  {object}  dtos.RaveResponse
// @Failure      500  {object}  dtos.RaveResponse
// @Security Bearer
// @Router       /api/v1/series/{id} [get]
func (seriesController *SeriesController) GetSeriesById(ctx *gin.Context) {
	calendarService := seriesController.SeriesService
	id, err := extractParamFromRequest("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
		return
	}
	calendar, err := calendarService.GetCalendar(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
		return
	}

	ctx.JSON(http.StatusCreated, &response.RaveResponse[*response.CreateCalendarResponse]{Data: calendar})

}
