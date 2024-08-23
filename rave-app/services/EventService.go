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
	log.Println("event request--->", createEventRequest)
	event := mapCreateEventRequestToEvent(createEventRequest)
	var calendar *models.Series
	var err error
	calendarService := NewSeriesService()
	if createEventRequest.CalendarId == 0 {
		calendar, err = calendarService.GetPublicCalendarFor(createEventRequest.OrganizerId)
		if err != nil {
			log.Println("error finding public calendar: ", err)
			return nil, err
		}
		log.Println("found public calendar: ", calendar)
	} else {
		calendar, err = calendarService.GetById(createEventRequest.CalendarId)
		if err != nil {
			return nil, err
		}
	}
	log.Println("calendar--->", calendar)
	event.SeriesID = calendar.ID
	log.Println("event: ", event)
	savedEvent, err := eventRepository.Save(event)
	if err != nil {
		log.Println("err saving event: ", err)
		return nil, err
	}
	_, err = calendarService.AddEventToCalendar(calendar.ID, savedEvent)
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

func (raveEventService *raveEventService) GetAllEventsFor(calendarId uint64) ([]*models.Event, error) {
	events, err := eventRepository.FindAllByCalendar(calendarId)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func mapEventToEventResponse(event *models.Event) *response.EventResponse {
	calendarService := NewSeriesService()
	calendar, err := calendarService.GetById(event.SeriesID)
	if err != nil {
		return nil
	}
	return &response.EventResponse{
		ID:                 event.ID,
		Message:            "event created successfully",
		Name:               event.Name,
		Location:           event.Location,
		Date:               event.EventDate,
		Time:               event.StartTime,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		Calendar:           calendar.Name,
	}
}

func mapCreateEventRequestToEvent(createEventRequest *request.CreateEventRequest) *models.Event {
	return &models.Event{
		Name:               createEventRequest.Name,
		Location:           createEventRequest.Location,
		EventDate:          createEventRequest.Date,
		StartTime:          createEventRequest.Time,
		SeriesID:           createEventRequest.CalendarId,
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.NOT_STARTED,
	}
}
