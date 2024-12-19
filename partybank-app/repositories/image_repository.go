package repositories

import (
	"github.com/djfemz/organizer-service/partybank-app/models"
	"gorm.io/gorm"
)

type ImageRepository interface {
	crudRepository[models.Image, uint64]
	FindByImageId(imageId string) (*models.Image, error)
}

type partybankImageRepository struct {
	*repositoryImpl[models.Image, uint64]
}

func NewPartybankImageRepository(db *gorm.DB) ImageRepository {
	return &partybankImageRepository{
		&repositoryImpl[models.Image, uint64]{
			db,
		},
	}
}

func (imageRepository partybankImageRepository) FindByImageId(imageId string) (*models.Image, error) {
	image := &models.Image{}
	if err := imageRepository.Db.Where(&models.Image{ImageId: imageId}).First(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}
