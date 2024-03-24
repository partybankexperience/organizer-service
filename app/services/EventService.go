package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*response.RaveResponse[any], error)
}

type raveEventService struct {
	repositories.EventRepository
	OrganizerService
}

func NewEventService() EventService {
	return &raveEventService{
		repositories.NewEventRepository(),
		NewOrganizerService(),
	}
}

func (raveEventService *raveEventService) Create(createEventRequest *request.CreateEventRequest) (*response.RaveResponse[any], error) {
	organizerService := raveEventService.OrganizerService
	organizer, err := organizerService.GetById(createEventRequest.OrganizerId)
	if err != nil {
		return nil, err
	}
	event := mapCreateEventRequestToEvent(createEventRequest)
	event.Organizer = organizer

	return nil, nil
}

func mapCreateEventRequestToEvent(createEventRequest *request.CreateEventRequest) *models.Event {
	return &models.Event{
		Name:               createEventRequest.Name,
		Location:           createEventRequest.Location,
		Date:               createEventRequest.Date,
		Time:               createEventRequest.Time,
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.NOT_STARTED,
	}
}
