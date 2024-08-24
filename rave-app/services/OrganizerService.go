package services

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/security/otp"

	"log"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateUserRequest) (*response.CreateOrganizerResponse, error)
	GetByUsername(username string) (*models.Organizer, error)
	UpdateOtpFor(id uint64, testOtp *otp.OneTimePassword) (*models.Organizer, error)
	GetById(id uint64) (*models.Organizer, error)
	AddEventStaff(staff *request.AddEventStaffRequest) (*response.RaveResponse[string], error)
	GetByOtp(otp string) (*models.Organizer, error)
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
	calendarService := NewSeriesService()

	organizer.Otp = password
	savedOrganizer, err := organizerService.Repository.Save(organizer)
	createCalendarRequest := &request.CreateSeriesRequest{
		Name:        "Public",
		OrganizerID: savedOrganizer.ID,
	}
	calendarResponse, err := calendarService.AddSeries(createCalendarRequest)
	if err != nil {
		return nil, err
	}
	calendar, err := calendarService.GetById(calendarResponse.ID)

	if err != nil {
		return nil, err
	}
	savedOrganizer.Series = append(savedOrganizer.Series, calendar)
	savedOrganizer, err = organizerService.Repository.Save(savedOrganizer)

	go func() {
		mailService.Send(request.NewEmailNotificationRequest(CreateNewOrganizerEmail(password.Code), organizer.Username))
	}()
	if savedOrganizer != nil {
		return &response.CreateOrganizerResponse{
			Message:  response.USER_CREATED_SUCCESSFULLY,
			Username: savedOrganizer.Username,
		}, nil
	}
	return nil, err
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
	organizer, err := organizerRepository.FindById(id)
	if organizer != nil {
		organizer.Otp = otp
		organizer, err = organizerRepository.Save(organizer)
		if err != nil {
			return nil, err
		}
		return organizer, nil
	} else {
		return nil, err
	}
}

func (organizerService *appOrganizerService) GetById(id uint64) (*models.Organizer, error) {
	organizationRepository := organizerService.Repository
	org, err := organizationRepository.FindById(id)
	if org == nil {
		return nil, err
	}
	return org, nil
}

func (organizerService *appOrganizerService) AddEventStaff(addStaffRequest *request.AddEventStaffRequest) (*response.RaveResponse[string], error) {
	res, err := organizerService.eventStaffService.Create(&request.CreateEventStaffRequest{StaffEmails: addStaffRequest.StaffEmails, EventId: addStaffRequest.EventId})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (organizerService *appOrganizerService) GetByOtp(otp string) (*models.Organizer, error) {
	organizerRepository := organizerService.Repository
	log.Println("repo:", organizerRepository)
	return organizerRepository.FindByOtp(otp)
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
