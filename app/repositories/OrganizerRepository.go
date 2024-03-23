package repositories

import (
	"github.com/djfemz/rave/app/models"
)

type OrganizerRepository interface {
	crudRepository[models.Organizer, uint64]
	FindByUsername(username string) (*models.Organizer, error)
}

type organizerRepositoryImpl struct {
	repositoryImpl[models.Organizer, uint64]
}

func NewOrganizerRepository() OrganizerRepository {
	var organizerRepository OrganizerRepository = &organizerRepositoryImpl{}
	return organizerRepository
}

func (organizer *organizerRepositoryImpl) FindByUsername(username string) (*models.Organizer, error) {
	db = connect()
	var organization = new(models.Organizer)
	err := db.Where("username=?", username).First(&organization).Error
	return organization, err
}
