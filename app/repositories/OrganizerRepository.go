package repositories

import (
	"github.com/djfemz/rave/app/models"
	otp2 "github.com/djfemz/rave/app/security/otp"
)

type OrganizerRepository interface {
	crudRepository[models.Organizer, uint64]
	FindByUsername(username string) (*models.Organizer, error)
	FindByOtp(otp string) (*models.Organizer, error)
}

type organizerRepositoryImpl struct {
	repositoryImpl[models.Organizer, uint64]
}

func NewOrganizerRepository() OrganizerRepository {
	var organizerRepository OrganizerRepository = &organizerRepositoryImpl{
		repositoryImpl[models.Organizer, uint64]{},
	}
	return organizerRepository
}

func (organizerRepository *organizerRepositoryImpl) FindByUsername(username string) (*models.Organizer, error) {
	db = connect()
	var organization = new(models.Organizer)
	err := db.Where("username=?", username).First(&organization).Error
	return organization, err
}

func (organizerRepository *organizerRepositoryImpl) FindByOtp(otp string) (*models.Organizer, error) {
	var organizer models.Organizer
	db = connect()
	err := db.Where(&models.Organizer{Otp: &otp2.OneTimePassword{Code: otp}}).Find(&organizer).Error
	if err != nil {
		return nil, err
	}
	return &organizer, nil
}
