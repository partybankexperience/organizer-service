package services

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"

	dtos "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/djfemz/organizer-service/partybank-app/utils"
)

type FileUploadService interface {
	UploadImage(file multipart.File) (*dtos.ImageUploadResponse, error)
	GetImage(imageId string) ([]byte, error)
}

type PartybankFileUploadService struct {
	ImageService
}

func NewFileUploadService(imageService ImageService) *PartybankFileUploadService {
	return &PartybankFileUploadService{
		ImageService: imageService,
	}
}

func (partybankFileUploadService *PartybankFileUploadService) UploadImage(file multipart.File) (*dtos.ImageUploadResponse, error) {
	bs, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read from image file")
	}
	imageId := utils.GenerateImageId(6)
	//"https://organizer-service.onrender.com/"
	imageUrl := os.Getenv("APPLICATION_BASE_URL") + "/api/v1/image/" + imageId
	image := &models.Image{
		ImageId: imageId,
		Content: bs,
		Url:     imageUrl,
	}
	imageResponse, err := partybankFileUploadService.ImageService.AddImage(image)
	if err != nil {
		log.Println("Error: ", err)
		return nil, errors.New("failed to add image")
	}
	return imageResponse, nil
}

func (partybankFileUploadService *PartybankFileUploadService) GetImage(imageId string) ([]byte, error) {
	image, err := partybankFileUploadService.ImageService.GetImage(imageId)
	if err != nil {
		log.Println("Error: ", err)
		return nil, errors.New("failed to retrieve image")
	}
	return image.Content, nil
}
