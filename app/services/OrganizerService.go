package services

import (
	"errors"
	"fmt"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
	"github.com/djfemz/rave/app/security/otp"
	"log"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error)
	GetByUsername(username string) (*models.Organizer, error)
	UpdateOtpFor(id uint64, testOtp *otp.OneTimePassword) (*models.Organizer, error)
}

type appOrganizerService struct {
	Repository  repositories.OrganizerRepository
	mailService MailService
}

func NewOrganizerService() OrganizerService {
	return &appOrganizerService{
		Repository:  repositories.NewOrganizerRepository(),
		mailService: NewMailService(),
	}
}

func (organizerService *appOrganizerService) Create(createOrganizerRequest *request.CreateOrganizerRequest) (*response.CreateOrganizerResponse, error) {
	organizer := mapCreateOrganizerRequestTo(createOrganizerRequest)
	password := otp.GenerateOtp()
	log.Println("organizer: ", organizer.User, " password: ", password, "username: ", organizer.Username)
	organizerService.mailService.Send(request.NewEmailNotificationRequest(CreateNewOrganizerEmail(password.Code), organizer.Username))
	organizer.Otp = password
	savedOrganizer := organizerService.Repository.Save(organizer)
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
	return organizer, err
}

func (organizerService *appOrganizerService) UpdateOtpFor(id uint64, otp *otp.OneTimePassword) (*models.Organizer, error) {
	organizerRepository := organizerService.Repository
	organizer := organizerRepository.FindById(id)
	if organizer != nil {
		organizer.Otp = otp
		organizer = organizerRepository.Save(organizer)
		return organizer, nil
	} else {
		return nil, errors.New("organizer not found")
	}
}

func mapCreateOrganizerRequestTo(organizerRequest *request.CreateOrganizerRequest) *models.Organizer {
	log.Println("organizerRequest", organizerRequest)
	return &models.Organizer{
		User: &models.User{
			Username: organizerRequest.Username,
			Role:     models.ORGANIZER,
		},
	}
}

func CreateNewOrganizerEmail(otp string) string {
	return fmt.Sprintf("Your One Time Password is %s", otp)
}
