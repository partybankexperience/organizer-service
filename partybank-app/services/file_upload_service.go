package services

import (
	"errors"
	"io"
	"mime/multipart"

	dtos "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/djfemz/organizer-service/partybank-app/utils"
)


type FileUploadService interface {
	UploadImage(file multipart.File) (*dtos.ImageUploadResponse, error)
	GetImage(imageId string) (*dtos.ImageUploadResponse, error)
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
	imageId:=utils.GenerateImageId(6)
	imageUrl:="https://organizer-service.onrender.com/"+imageId
	image:=&models.Image{
		ImageId:      imageId,
		Content: string(bs),
        Url:    imageUrl,
	}
	imageResponse, err:=partybankFileUploadService.ImageService.AddImage(image)
	if err != nil {
		return nil, errors.New("failed to add image")
	}
	return imageResponse, nil
}

func (partybankFileUploadService *PartybankFileUploadService) GetImage(imageId string) (*dtos.ImageUploadResponse, error) {
	return partybankFileUploadService.ImageService.GetImage(imageId)
}

