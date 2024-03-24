package services

import (
	"errors"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
	"github.com/djfemz/rave/app/security/otp"

	"log"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateUserRequest) (*response.CreateOrganizerResponse, error)
	GetByUsername(username string) (*models.Organizer, error)
	UpdateOtpFor(id uint64, testOtp *otp.OneTimePassword) (*models.Organizer, error)
	GetById(id uint64) (*models.Organizer, error)
	AddEventStaff(staff *request.AddEventStaffRequest) (*response.RaveResponse[string], error)
	AddEvent(eventRequest *request.CreateEventRequest) (*response.RaveResponse[*response.EventResponse], error)
}

type appOrganizerService struct {
	Repository        repositories.OrganizerRepository
	eventStaffService EventStaffService
}

func NewOrganizerService() OrganizerService {
	return &appOrganizerService{
		Repository:        repositories.NewOrganizerRepository(),
		eventStaffService: NewEventStaffService(),
	}
}

func (organizerService *appOrganizerService) Create(createOrganizerRequest *request.CreateUserRequest) (*response.CreateOrganizerResponse, error) {
	organizer := mapCreateOrganizerRequestTo(createOrganizerRequest)
	password := otp.GenerateOtp()
	mailService := NewMailService()
	mailService.Send(request.NewEmailNotificationRequest(CreateNewOrganizerEmail(password.Code), organizer.Username))
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

func (organizerService *appOrganizerService) GetById(id uint64) (*models.Organizer, error) {
	organizationRepository := organizerService.Repository
	org := organizationRepository.FindById(id)
	if org == nil {
		return nil, errors.New("organization not found")
	}
	return org, nil
}

func (organizerService *appOrganizerService) AddEventStaff(addStaffRequest *request.AddEventStaffRequest) (*response.RaveResponse[string], error) {
	res, err := organizerService.eventStaffService.Create(&request.CreateEventStaffRequest{StaffEmails: addStaffRequest.StaffEmails})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (organizerService *appOrganizerService) AddEvent(eventRequest *request.CreateEventRequest) (*response.RaveResponse[*response.EventResponse], error) {
	eventService := NewEventService()
	return eventService.Create(eventRequest)
}

func mapCreateOrganizerRequestTo(organizerRequest *request.CreateUserRequest) *models.Organizer {
	log.Println("organizerRequest", organizerRequest)
	return &models.Organizer{
		User: &models.User{
			Username: organizerRequest.Username,
			Role:     models.ORGANIZER,
		},
	}
}

func CreateNewOrganizerEmail(content string) string {
	return content
}
