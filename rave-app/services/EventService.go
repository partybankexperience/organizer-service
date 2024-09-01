package services

import (
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*response.EventResponse, error)
	GetById(id uint64) (*response.EventResponse, error)
	GetEventBy(id uint64) (*models.Event, error)
	UpdateEventInformation(id uint64, updateRequest *request.UpdateEventRequest) (*response.EventResponse, error)
	UpdateEvent(event *models.Event) error
	GetAllEventsFor(organizerId uint64, pageNumber int, pageSize int) ([]*response.EventResponse, error)
}

type raveEventService struct {
	repositories.EventRepository
	OrganizerService
	SeriesService
}

func NewEventService(eventRepository repositories.EventRepository, organizerService OrganizerService, seriesService SeriesService) EventService {
	return &raveEventService{
		eventRepository,
		organizerService,
		seriesService,
	}
}

func (raveEventService *raveEventService) Create(createEventRequest *request.CreateEventRequest) (*response.EventResponse, error) {
	event := mapCreateEventRequestToEvent(createEventRequest)
	var calendar *models.Series
	var err error

	if createEventRequest.SeriesId == 0 {
		calendar, err = raveEventService.GetPublicCalendarFor(createEventRequest.OrganizerId)
		if err != nil {
			log.Println("error finding public calendar: ", err)
			return nil, err
		}
		log.Println("found public calendar: ", calendar)
	} else {
		calendar, err = raveEventService.SeriesService.GetById(createEventRequest.SeriesId)
		if err != nil {
			return nil, err
		}
	}
	event.SeriesID = calendar.ID
	savedEvent, err := raveEventService.Save(event)
	if err != nil {
		log.Println("err saving event: ", err)
		return nil, err
	}
	_, err = raveEventService.AddEventToCalendar(calendar.ID, savedEvent)
	if err != nil {
		return nil, err
	}
	res := mapEventToEventResponse(savedEvent, raveEventService.SeriesService)
	return res, nil
}

func (raveEventService *raveEventService) GetById(id uint64) (*response.EventResponse, error) {
	foundEvent, err := raveEventService.FindById(id)
	if err != nil {
		return nil, err
	}
	return mapEventToEventResponse(foundEvent, raveEventService.SeriesService), nil
}

func (raveEventService *raveEventService) GetEventBy(id uint64) (*models.Event, error) {
	return raveEventService.FindById(id)
}

func (raveEventService *raveEventService) UpdateEventInformation(id uint64, updateRequest *request.UpdateEventRequest) (*response.EventResponse, error) {
	updateEventResponse := &response.EventResponse{}

	foundEvent, err := raveEventService.FindById(id)
	if err != nil {
		return nil, err
	}
	copyErrors := model.Copy(foundEvent, updateRequest)
	if len(copyErrors) != 0 {
		return nil, errors.New("could not update event")
	}
	savedEvent, err := raveEventService.Save(foundEvent)
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
	_, err := raveEventService.Save(event)
	return err
}

func (raveEventService *raveEventService) GetAllEventsFor(calendarId uint64, pageNumber, pageSize int) ([]*response.EventResponse, error) {
	events, err := raveEventService.FindAllByCalendar(calendarId, pageNumber, pageSize)
	eventsResponses := make([]*response.EventResponse, 0)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		eventResponse := mapEventToEventResponse(event, raveEventService.SeriesService)
		eventsResponses = append(eventsResponses, eventResponse)
	}
	return eventsResponses, nil
}

func mapEventToEventResponse(event *models.Event, seriesService SeriesService) *response.EventResponse {
	series, err := seriesService.GetById(event.SeriesID)
	if err != nil {
		return nil
	}
	eventResponse := &response.EventResponse{
		ID:      event.ID,
		Message: "event created successfully",
		Name:    event.Name,

		Date:               event.EventDate,
		Time:               event.StartTime,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		SeriesID:           series.ID,
		Venue:              event.Venue,
		MapUrl:             event.MapUrl,
		MapEmbeddedUrl:     event.MapEmbeddedUrl,
		AttendeeTerm:       event.AttendeeTerm,
		EventTheme:         event.EventTheme,
	}

	if event.Location != nil {
		eventResponse.Location = &models.Location{
			State:   event.Location.State,
			Country: event.Location.Country,
		}
	}
	return eventResponse
}

func mapCreateEventRequestToEvent(createEventRequest *request.CreateEventRequest) *models.Event {
	return &models.Event{
		Name: createEventRequest.Name,
		Location: &models.Location{
			State:   createEventRequest.State,
			Country: createEventRequest.Country,
		},
		EventDate:          createEventRequest.Date,
		StartTime:          createEventRequest.Time,
		SeriesID:           createEventRequest.SeriesId,
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.NOT_STARTED,
		MapUrl:             createEventRequest.MapUrl,
		MapEmbeddedUrl:     createEventRequest.MapEmbeddedUrl,
		EventTheme:         createEventRequest.EventTheme,
		AttendeeTerm:       createEventRequest.AttendeeTerm,
		Venue:              createEventRequest.Venue,
	}
}
