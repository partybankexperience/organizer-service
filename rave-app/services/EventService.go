package services

import (
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/mappers"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/utils"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (*response.EventResponse, error)
	GetById(id uint64) (*response.EventResponse, error)
	GetEventBy(id uint64) (*models.Event, error)
	GetEventByReference(reference string) (*response.EventResponse, error)
	DiscoverEvents(page int, size int) ([]*response.EventResponse, error)
	UpdateEventInformation(id uint64, updateRequest *request.UpdateEventRequest) (*response.EventResponse, error)
	UpdateEvent(event *models.Event) error
	GetAllEventsFor(organizerId uint64, pageNumber int, pageSize int) ([]*response.EventResponse, error)
	PublishEvent(eventId uint64) (*response.EventResponse, error)
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
	org, err := raveEventService.OrganizerService.GetById(createEventRequest.OrganizerId)
	if err != nil || org == nil {
		log.Println("err finding organizer: ", err)
		return nil, err
	}
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
	log.Println("organizer: ", *org)
	event.SeriesID = calendar.ID
	event.CreatedBy = calendar.Name
	event.PublicationState = models.DRAFT
	savedEvent, err := raveEventService.Save(event)
	if err != nil {
		log.Println("err saving event: ", err)
		return nil, err
	}
	_, err = raveEventService.AddEventToCalendar(calendar.ID, savedEvent)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	res := mappers.MapEventToEventResponse(savedEvent)
	return res, nil
}

func (raveEventService *raveEventService) GetById(id uint64) (*response.EventResponse, error) {
	foundEvent, err := raveEventService.FindById(id)
	if err != nil {
		return nil, err
	}
	log.Println("event: ", *foundEvent)
	return mappers.MapEventToEventResponse(foundEvent), nil
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
		eventResponse := mappers.MapEventToEventResponse(event)
		eventsResponses = append(eventsResponses, eventResponse)
	}
	return eventsResponses, nil
}

func (raveEventService *raveEventService) DiscoverEvents(page int, size int) ([]*response.EventResponse, error) {
	events, err := raveEventService.FindAllPublishedByPage(page, size)
	if err != nil {
		return nil, err
	}
	log.Println("events: ", events)

	allEvents := mappers.MapEventsToEventResponses(events)
	return allEvents, nil
}

func (raveEventService *raveEventService) GetEventByReference(reference string) (*response.EventResponse, error) {
	event, err := raveEventService.EventRepository.FindByReference(reference)
	if err != nil {
		return nil, errors.New("failed to find requested event")
	}
	eventResponse := mappers.MapEventToEventResponse(event)
	return eventResponse, nil
}

func (raveEventService *raveEventService) PublishEvent(eventId uint64) (*response.EventResponse, error) {
	event, err := raveEventService.GetEventBy(eventId)
	if err != nil {
		return nil, errors.New("event not found")
	}
	if event.Tickets != nil && len(event.Tickets) > 0 {
		event.PublicationState = models.PUBLISHED
	}
	event, err = raveEventService.Save(event)
	if err != nil {
		return nil, errors.New("failed to save event")
	}
	return mappers.MapEventToEventResponse(event), nil
}

func mapCreateEventRequestToEvent(createEventRequest *request.CreateEventRequest) *models.Event {
	return &models.Event{
		Name: createEventRequest.Name,
		Location: &models.Location{
			State:   createEventRequest.State,
			Country: createEventRequest.Country,
			City:    createEventRequest.City,
		},
		EventDate:          createEventRequest.Date,
		StartTime:          createEventRequest.Time,
		SeriesID:           createEventRequest.SeriesId,
		ContactInformation: createEventRequest.ContactInformation,
		Description:        createEventRequest.Description,
		Status:             models.UPCOMING,
		MapUrl:             createEventRequest.MapUrl,
		MapEmbeddedUrl:     createEventRequest.MapEmbeddedUrl,
		EventTheme:         createEventRequest.EventTheme,
		AttendeeTerm:       createEventRequest.AttendeeTerm,
		Venue:              createEventRequest.Venue,
		ImageUrl:           createEventRequest.ImageUrl,
		Reference:          utils.GenerateEventReference(),
	}
}
