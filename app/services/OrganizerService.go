package services

import (
	"errors"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error)
	GetByUsername(username string) (*models.Organizer, error)
}

type AppOrganizerService struct {
	repository repositories.Repository[models.Organizer, uint64]
}

func (organizerService *AppOrganizerService) Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error) {
	savedOrganizer := organizerService.repository.Save(mapCreateOrganizerRequestTo(createOrganizerRequest))
	if savedOrganizer != nil {
		return &response.CreateOrganizerResponse{
			Message:  response.USER_CREATED_SUCCESSFULLY,
			Username: savedOrganizer.Username,
		}, nil
	}
	return nil, errors.New("failed to create user with username")
}

func (organizerService *AppOrganizerService) GetByUsername(username string) (*models.Organizer, error) {
	organizer := organizerService.repository.FindByUsername(username)
	return organizer, nil
}

func mapCreateOrganizerRequestTo(organizerRequest *request.CreateOrganizerRequest) *models.Organizer {
	return &models.Organizer{
		User: &models.User{
			Username: organizerRequest.Username,
			Role:     models.ORGANIZER,
		},
	}
}
