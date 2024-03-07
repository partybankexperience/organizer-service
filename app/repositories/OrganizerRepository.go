package repositories

import (

	"github.com/djfemz/rave/app/models"
)

type OrganizerRepository interface {
	Save(organizer *models.Organizer) (*models.Organizer, error)
	FindById(id uint) *models.Organizer
}

type OrganizerRepositoryImpl struct{
}


func NewOrganizerRepositoryImpl() *OrganizerRepositoryImpl{
	return &OrganizerRepositoryImpl{}
}

func (org *OrganizerRepositoryImpl) Save(organizer *models.Organizer) (*models.Organizer, error){

	return nil, nil
}

func (org *OrganizerRepositoryImpl) FindById(id uint) *models.Organizer{
	return nil
}




