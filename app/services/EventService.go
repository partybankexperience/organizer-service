package services

import (
	"errors"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*response.RaveResponse[*response.EventResponse], error)
	GetById(i uint64) (*response.EventResponse, error)
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

func (raveEventService *raveEventService) Create(createEventRequest *request.CreateEventRequest) (*response.RaveResponse[*response.EventResponse], error) {
	organizerService := raveEventService.OrganizerService
	organizer, err := organizerService.GetById(createEventRequest.OrganizerId)
	if err != nil {
		return nil, err
	}
	event := mapCreateEventRequestToEvent(createEventRequest)
	event.Organizer = organizer
	eventRepository := raveEventService.EventRepository
	savedEvent := eventRepository.Save(event)
	if savedEvent == nil {
		return nil, err
	}
	return &response.RaveResponse[*response.EventResponse]{
		Data: mapEventToEventResponse(savedEvent),
	}, nil
}

func (raveEventService *raveEventService) GetById(id uint64) (*response.EventResponse, error) {
	foundEvent := raveEventService.EventRepository.FindById(id)
	if foundEvent == nil {
		return nil, errors.New("event not found")
	}
	return mapEventToEventResponse(foundEvent), nil
}

func mapEventToEventResponse(event *models.Event) *response.EventResponse {
	return &response.EventResponse{
		Message:            "event created successfully",
		Name:               event.Name,
		Organizer:          event.Organizer.Name,
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
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.NOT_STARTED,
	}
}
