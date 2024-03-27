package services

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
)

type EventStaffService interface {
	Create(createUserRequest *request.CreateEventStaffRequest) (*response.RaveResponse[string], error)
}

type raveEventStaffService struct {
	EventService
}

var eventStaffRepository = repositories.NewEventStaffRepository()

func NewEventStaffService() EventStaffService {
	return &raveEventStaffService{
		NewEventService(),
	}
}

func (eventStaffService *raveEventStaffService) Create(createUserRequest *request.CreateEventStaffRequest) (*response.RaveResponse[string], error) {
	eventService := NewEventService()
	event, err := eventService.GetEventBy(createUserRequest.EventId)
	if err != nil {
		return nil, err
	}
	mailService := NewMailService()
	for _, email := range createUserRequest.StaffEmails {
		savedStaff, err := updateEvent(createUserRequest, email, event, eventStaffRepository)
		if err != nil {
			return nil, err
		} else if savedStaff != nil {
			notification := request.NewEmailNotificationRequest(email, `
				welcome to rave, sign in using this email address
			`)
			_, err := mailService.Send(notification)
			if err != nil {
				return nil, err
			}
		}
	}

	return &response.RaveResponse[string]{Data: "event staffs invited"}, nil
}

func updateEvent(createUserRequest *request.CreateEventStaffRequest, email string, event *models.Event, repo repositories.EventStaffRepository) (*models.EventStaff, error) {
	eventStaff := mapMailToEventStaff(email)
	eventStaff.EventID = event.ID
	eventStaff.EventID = createUserRequest.EventId
	event.EventStaff = append(event.EventStaff, eventStaff)
	savedStaff, err := repo.Save(eventStaff)
	return savedStaff, err
}

func mapMailToEventStaff(email string) *models.EventStaff {
	return &models.EventStaff{
		User: &models.User{
			Username: email,
			Role:     models.EVENT_STAFF,
		},
	}
}
