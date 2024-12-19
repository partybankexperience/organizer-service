package repositories

import "github.com/djfemz/organizer-service/partybank-app/models"

type ImageRepository interface {
	crudRepository[models.Image, uint64]
	FindByImageId(imageId string) (*models.Image, error)
}

type partybankImageRepository struct {
	repositoryImpl[models.Image, uint64]
}

func NewPartybankImageRepository() ImageRepository {
    return &partybankImageRepository{}
}

func (imageRepository partybankImageRepository) FindByImageId(imageId string) (*models.Image, error) {
	image:=&models.Image{}
	if err:=imageRepository.Db.Where(&models.Image{ImageId: imageId}).First(&image).Error;err!=nil{
		return nil, err
	}
	return image, nil
}

