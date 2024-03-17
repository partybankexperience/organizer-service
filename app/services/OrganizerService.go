package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateOrganizerRequest) *response.CreateOrganizerResponse
	GetByUsername(username string) *models.Organizer
}

type AppOrganizerService struct {
}

func (organizerService *AppOrganizerService) Create(createOrganizerRequest *request.CreateOrganizerRequest) *response.CreateOrganizerResponse {
	return nil
}

func (organizerService *AppOrganizerService) GetByUsername(username string) *models.Organizer {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &models.Organizer{
		ID:       1,
		Username: "jon@email.com",
		Password: string(hashPassword),
		Role:     models.ORGANIZER,
	}
}
