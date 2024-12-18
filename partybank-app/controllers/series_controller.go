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
)

type SeriesController struct {
	services.SeriesService
	*validator.Validate
}

func NewSeriesController(seriesService services.SeriesService, objectValidator *validator.Validate) *SeriesController {
	return &SeriesController{
		seriesService,
		objectValidator,
	}
}

// CreateSeries godoc
// @Summary      Add New Series
// @Description  Adds New Series
// @Tags         Series
// @Accept       json
// @Param 		 tags body dtos.CreateSeriesRequest true "Series tags"
// @Produce      json
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/series [post]
func (seriesController *SeriesController) CreateSeries(ctx *gin.Context) {
	createSeriesRequest := &request.CreateSeriesRequest{}
	err := ctx.BindJSON(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	err = seriesController.Struct(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	resp, err := seriesController.AddSeries(createSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	ctx.JSON(http.StatusCreated, &response.PartybankBaseResponse[*response.CreateCalendarResponse]{Data: resp})
}

// GetSeriesById godoc
// @Summary      Get Series by id
// @Description  Get Series by id
// @Tags         Series
// @Accept       json
// @Produce      json
// @Param        id   path   int  true  "series id"
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/series/{id} [get]
func (seriesController *SeriesController) GetSeriesById(ctx *gin.Context) {
	seriesService := seriesController.SeriesService
	id, err := extractIdParamFromRequest("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	calendar, err := seriesService.GetCalendar(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}

	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[*response.SeriesResponse]{Data: calendar})

}

// GetSeriesForOrganizer godoc
// @Summary      Get Series by organizerId
// @Description  Get Series by organizerId
// @Tags         Series
// @Accept       json
// @Produce      json
// @Param        organizerId   path   int  true  "organizerId"
// @Param        page   query   int  true  "page"
// @Param        size   query   int  true  "size"
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/series/organizer/{organizerId} [get]
func (seriesController *SeriesController) GetSeriesForOrganizer(ctx *gin.Context) {
	organizerId, err := extractIdParamFromRequest("organizerId", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	log.Println("org id", organizerId)
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageNumber, err := utils.ConvertQueryStringToInt(page)
	if err != nil {
		handleError(ctx, err)
		return
	}
	pageSize, err := utils.ConvertQueryStringToInt(size)
	if err != nil {
		handleError(ctx, err)
		return
	}
	orgSeries, err := seriesController.SeriesService.GetSeriesFor(organizerId, pageNumber, pageSize)
	if err != nil {
		handleError(ctx, err)
		return
	}
	log.Println("series for org: ", orgSeries)
	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[[]*response.SeriesResponse]{Data: orgSeries})
}

// UpdateSeries godoc
// @Summary      Update Existing Series
// @Description  Update Existing Series
// @Tags         Series
// @Accept       json
// @Param 		 tags body dtos.UpdateSeriesRequest true "Series tags"
// @Param        seriesId   path   int  true  "seriesId"
// @Produce      json
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/series [put]
func (seriesController *SeriesController) UpdateSeries(ctx *gin.Context) {
	seriesId, err := extractIdParamFromRequest("seriesId", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	updateSeriesRequest := &request.UpdateSeriesRequest{}
	err = ctx.BindJSON(updateSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	err = seriesController.Struct(updateSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	resp, err := seriesController.SeriesService.UpdateSeries(seriesId, updateSeriesRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[*response.SeriesResponse]{Data: resp})
}

// AddEventToSeries godoc
// @Summary      Adds event to series
// @Description  Changes the series an event belongs to
// @Tags         Series
// @Accept       json
// @Produce      json
// @Param        seriesId   path   int  true  "seriesId"
// @Param        eventId   query   int  true  "eventId"
// @Success      200  {object}  dtos.PartybankBaseResponse
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Security Bearer
// @Router       /api/v1/series/events/add [get]
func (seriesController *SeriesController) AddEventToSeries(ctx *gin.Context) {
	seriesId, err := extractIdParamFromRequest("seriesId", ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	eventId, err := utils.ConvertQueryStringToInt(ctx.Query("eventId"))
	if err != nil {
		handleError(ctx, errors.New("failed to get eventId from request"))
		return
	}
	resp, err := seriesController.SeriesService.AddToSeries(seriesId, uint64(eventId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.PartybankBaseResponse[error]{Data: err})
		return
	}
	ctx.JSON(http.StatusOK, &response.PartybankBaseResponse[*response.SeriesResponse]{Data: resp})
}
