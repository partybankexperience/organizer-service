package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error)
	GetByUsername(username string) (*models.Organizer, error)
}

type AppOrganizerService struct {
}

func (organizerService *AppOrganizerService) Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error) {
	return nil, nil
}

func (organizerService *AppOrganizerService) GetByUsername(username string) (*models.Organizer, error) {
	return nil, nil
}
