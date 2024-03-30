package services

import (
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*models.Event, error)
	GetById(id uint64) (*response.EventResponse, error)
	GetEventBy(id uint64) (*models.Event, error)
	UpdateEventInformation(id uint64, updateRequest *request.UpdateEventRequest) (*response.EventResponse, error)
	UpdateEvent(event *models.Event) error
	GetAllEventsFor(organizerId uint64) ([]*models.Event, error)
}

var eventRepository = repositories.NewEventRepository()

type raveEventService struct {
	OrganizerService
}

func NewEventService() EventService {
	return &raveEventService{}
}

func (raveEventService *raveEventService) Create(createEventRequest *request.CreateEventRequest) (*models.Event, error) {
	event := mapCreateEventRequestToEvent(createEventRequest)
	savedEvent, err := eventRepository.Save(event)
	if err != nil {
		return nil, err
	}
	return savedEvent, nil
}

func (raveEventService *raveEventService) GetById(id uint64) (*response.EventResponse, error) {
	foundEvent, err := eventRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return mapEventToEventResponse(foundEvent), nil
}

func (raveEventService *raveEventService) GetEventBy(id uint64) (*models.Event, error) {
	return eventRepository.FindById(id)
}

func (raveEventService *raveEventService) UpdateEventInformation(id uint64, updateRequest *request.UpdateEventRequest) (*response.EventResponse, error) {
	updateEventResponse := &response.EventResponse{}

	foundEvent, err := eventRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	copyErrors := model.Copy(foundEvent, updateRequest)
	if len(copyErrors) != 0 {
		return nil, errors.New("could not update event")
	}
	savedEvent, err := eventRepository.Save(foundEvent)
	if err != nil {
		return nil, err
	}
	copyErrors = model.Copy(updateEventResponse, savedEvent)
	if len(copyErrors) != 0 {
		return nil, errors.New("could not update event")
	}
	return updateEventResponse, nil
}

func (raveEventService *raveEventService) UpdateEvent(event *models.Event) error {
	_, err := eventRepository.Save(event)
	return err
}

func (raveEventService *raveEventService) GetAllEventsFor(organizerId uint64) ([]*models.Event, error) {
	events, err := eventRepository.FindAllByOrganizer(organizerId)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func mapEventToEventResponse(event *models.Event) *response.EventResponse {
	OrganizerService := NewOrganizerService()
	org, err := OrganizerService.GetById(event.OrganizerID)
	if err != nil {
		return nil
	}
	return &response.EventResponse{
		Message:            "event created successfully",
		Name:               event.Name,
		Location:           event.Location,
		Date:               event.Date,
		Time:               event.Time,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		Organizer:          org.Username,
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
