package services

import (
	"errors"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
)

type EventStaffService interface {
	Create(createUserRequest *request.CreateEventStaffRequest) (*response.RaveResponse[string], error)
}

type raveEventStaffService struct {
	repositories.EventStaffRepository
	EventService
	MailService
}

func NewEventStaffService() EventStaffService {
	return &raveEventStaffService{
		repositories.NewEventStaffRepository(),
		NewEventService(),
		NewMailService(),
	}
}

func (eventStaffService *raveEventStaffService) Create(createUserRequest *request.CreateEventStaffRequest) (*response.RaveResponse[string], error) {
	repo := eventStaffService.EventStaffRepository
	for _, email := range createUserRequest.StaffEmails {
		eventStaff := mapMailToEventStaff(email)
		savedStaff := repo.Save(eventStaff)
		if savedStaff != nil {
			notification := request.NewEmailNotificationRequest(email, `
				welcome to rave, sign in using this email address
			`)
			res, err := eventStaffService.MailService.Send(notification)
			if err != nil {
				return nil, err
			}
			return &response.RaveResponse[string]{Data: res}, nil
		}
	}

	return nil, errors.New("failed to add event staff")
}

func mapMailToEventStaff(email string) *models.EventStaff {
	return &models.EventStaff{
		User: &models.User{
			Username: email,
		},
	}
}
