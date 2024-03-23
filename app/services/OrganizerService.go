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

type appOrganizerService struct {
	Repository repositories.OrganizerRepository
}

func NewOrganizerService() OrganizerService {
	return &appOrganizerService{
		Repository: repositories.NewOrganizerRepository(),
	}
}

func (organizerService *appOrganizerService) Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error) {
	savedOrganizer := organizerService.Repository.Save(mapCreateOrganizerRequestTo(createOrganizerRequest))
	if savedOrganizer != nil {
		return &response.CreateOrganizerResponse{
			Message:  response.USER_CREATED_SUCCESSFULLY,
			Username: savedOrganizer.Username,
		}, nil
	}
	return nil, errors.New("failed to create user with username")
}

func (organizerService *appOrganizerService) GetByUsername(username string) (*models.Organizer, error) {
	organizer, err := organizerService.Repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
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
