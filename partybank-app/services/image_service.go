package services

import (
	"errors"
	"log"

	dtos "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
)

type ImageService interface {
	AddImage(image *models.Image) (*dtos.ImageUploadResponse, error)
	GetImage(imageId string) (*models.Image, error)
}

type PartybankImageService struct {
	repositories.ImageRepository
}

func NewImageService(imageRepository repositories.ImageRepository) ImageService {
	return &PartybankImageService{
		imageRepository,
	}
}

func (partybankImageService *PartybankImageService) AddImage(image *models.Image) (*dtos.ImageUploadResponse, error) {
	image, err := partybankImageService.ImageRepository.Save(image)
	if err != nil {
		log.Println("Error saving image: ", err)
		return nil, errors.New("failed to save image: " + err.Error())
	}
	return &dtos.ImageUploadResponse{
		Url: image.Url,
	}, nil
}

func (partybankImageService *PartybankImageService) GetImage(imageId string) (*models.Image, error) {
	image, err := partybankImageService.ImageRepository.FindByImageId(imageId)
	if err != nil {
		return nil, errors.New("failed to retrieve image: " + err.Error())
	}
	return image, nil
}
