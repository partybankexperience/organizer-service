package services

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*models.Event, error)
	GetById(id uint64) (*response.EventResponse, error)
	GetEventBy(id uint64) (*models.Event, error)
}

type raveEventService struct {
	OrganizerService
}

func NewEventService() EventService {
	return &raveEventService{}
}

func (raveEventService *raveEventService) Create(createEventRequest *request.CreateEventRequest) (*models.Event, error) {
	event := mapCreateEventRequestToEvent(createEventRequest)
	eventRepository := repositories.NewEventRepository()
	savedEvent, err := eventRepository.Save(event)
	if err != nil {
		return nil, err
	}
	return savedEvent, nil
}

func (raveEventService *raveEventService) GetById(id uint64) (*response.EventResponse, error) {
	foundEvent, err := repositories.NewEventRepository().FindById(id)
	if err != nil {
		return nil, err
	}
	return mapEventToEventResponse(foundEvent), nil
}

func (raveEventService *raveEventService) GetEventBy(id uint64) (*models.Event, error) {
	raveEventRepository := repositories.NewEventRepository()
	return raveEventRepository.FindById(id)
}

func mapEventToEventResponse(event *models.Event) *response.EventResponse {
	return &response.EventResponse{
		Message:            "event created successfully",
		Name:               event.Name,
		Location:           event.Location,
		Date:               event.Date,
		Time:               event.Time,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
	}
}

func mapCreateEventRequestToEvent(createEventRequest *request.CreateEventRequest) *models.Event {
	return &models.Event{
		Name:               createEventRequest.Name,
		Location:           createEventRequest.Location,
		Date:               createEventRequest.Date,
		Time:               createEventRequest.Time,
		OrganizerID:        createEventRequest.OrganizerId,
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.NOT_STARTED,
	}
}
