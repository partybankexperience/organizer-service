package controllers

import (
	"errors"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/services"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
)

type FileController struct {
	services.FileUploadService
}

func NewFileController(fileService services.FileUploadService) *FileController {
	return &FileController{
		fileService,
	}
}

// UploadImage godoc
// @Summary Upload an image
// @Description Upload an image file to the server
// @Tags Images
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image to upload"
// @Success 201 {object} dtos.PartybankBaseResponse "success"
// @Failure 400 {object} dtos.PartybankBaseResponse "error"
// @Router /api/v1/image [post]
func (fileController *FileController) UploadImage(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("could not find file in request"))
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("error closing file: ", err)
		}
	}(file)
	imageUploadResponse, err := fileController.FileUploadService.UploadImage(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &response.PartybankBaseResponse[*response.ImageUploadResponse]{Data: imageUploadResponse})
}

// GetImage godoc
// @Summary      Get Image
// @Description  Get Image
// @Tags         Images
// @Accept       json
// @Param        image_id  path string  true  "image_id"
// @Produce      application/octet-stream
// @Produce      json
// @Success      200 {string} binary "Binary data as a response"
// @Failure      400  {object}  dtos.PartybankBaseResponse
// @Router     /api/v1/image/{image_id} [get]
func (fileController *FileController) GetImage(ctx *gin.Context) {
	imageId := ctx.Param("image_id")
	imageResponse, err := fileController.FileUploadService.GetImage(imageId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.PartybankBaseResponse[string]{Data: err.Error()})
		return
	}
	ctx.Data(http.StatusOK, "image/png", imageResponse)
	return
}
