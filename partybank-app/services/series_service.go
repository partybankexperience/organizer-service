package services

import (
	"errors"
	dtos "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/mappers"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type SeriesService interface {
	AddSeries(createCalendarRequest *dtos.CreateSeriesRequest) (*response.CreateCalendarResponse, error)
	GetById(id uint64) (*models.Series, error)
	AddEventToSeries(id uint64, event *models.Event) (*models.Series, error)
	AddToSeries(seriesId, eventId uint64) (*response.SeriesResponse, error)
	GetCalendar(id uint64) (*response.SeriesResponse, error)
	GetPublicCalendarFor(id uint64) (*models.Series, error)
	GetSeriesFor(organizerId uint64, pageNumber int, pageSize int) ([]*response.SeriesResponse, error)
	GetSeriesOrganizer(seriesId uint64) (uint64, error)
	UpdateSeries(seriesId uint64, series *dtos.UpdateSeriesRequest) (*response.SeriesResponse, error)
	SetEventService(eventService EventService)
}

type raveSeriesService struct {
	repositories.SeriesRepository
	EventService
}

func NewSeriesService(seriesRepository repositories.SeriesRepository) SeriesService {
	return &raveSeriesService{seriesRepository, nil}
}

func (raveSeriesService *raveSeriesService) AddSeries(createSeriesRequest *dtos.CreateSeriesRequest) (*response.CreateCalendarResponse, error) {
	calendar := &models.Series{}
	errs := model.Copy(calendar, createSeriesRequest)
	log.Println("calendar: ", *calendar)
	isCopyErrorPresent := len(errs) > 0
	if isCopyErrorPresent {
		return nil, errors.New("error creating calendar from request")
	}
	savedCalendar, err := raveSeriesService.SeriesRepository.Save(calendar)
	log.Println("saved: ", *savedCalendar)
	if err != nil {
		return nil, err
	}
	createdCalendarResponse := &response.CreateCalendarResponse{}
	createdCalendarResponse.ID = savedCalendar.ID
	createdCalendarResponse.Name = savedCalendar.Name
	createdCalendarResponse.Message = "calendar created successfully"
	return createdCalendarResponse, nil
}

func (raveSeriesService *raveSeriesService) GetById(id uint64) (*models.Series, error) {
	calendar, err := raveSeriesService.SeriesRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (raveSeriesService *raveSeriesService) GetCalendar(id uint64) (*response.SeriesResponse, error) {
	calendar, err := raveSeriesService.GetById(id)
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	resp := mapSeriesToSeriesResponse(calendar)
	return resp, nil
}

func mapSeriesToSeriesResponse(series *models.Series) *response.SeriesResponse {
	events := make([]*response.EventResponse, 0)
	for _, event := range series.Events {
		seriesEvent := &response.EventResponse{
			ID:          event.ID,
			Name:        event.Name,
			Date:        event.EventDate,
			Time:        event.StartTime,
			Description: event.Description,
			Status:      event.Status,
		}
		events = append(events, seriesEvent)
	}
	res := &response.SeriesResponse{}
	res.Name = series.Name
	res.ImageUrl = series.ImageUrl
	res.Description = series.Description
	res.ID = series.ID
	res.OrganizerID = series.OrganizerID
	res.Events = events
	return res
}

func (raveSeriesService *raveSeriesService) AddEventToSeries(id uint64, event *models.Event) (*models.Series, error) {

	calendar, err := raveSeriesService.GetById(id)
	if err != nil {
		return nil, err
	}
	log.Println("event: ", *event)
	calendar.Events = append(calendar.Events, event)
	log.Println("calendar: ", calendar.Events)
	calendar, err = raveSeriesService.SeriesRepository.Save(calendar)
	if err != nil {
		return nil, errors.New("error Adding event to series")
	}
	return calendar, nil
}

func (raveSeriesService *raveSeriesService) GetPublicCalendarFor(id uint64) (*models.Series, error) {
	calendar, err := raveSeriesService.SeriesRepository.FindPublicSeriesFor(id)
	if err != nil {
		return nil, errors.New("error finding requested resource")
	}
	return calendar, nil
}

func (raveSeriesService *raveSeriesService) GetSeriesFor(organizerId uint64, pageNumber int, pageSize int) ([]*response.SeriesResponse, error) {
	userSeries, err := raveSeriesService.FindAllSeriesFor(organizerId, pageNumber, pageSize)
	if err != nil {
		log.Println("Error: ", err)
		return nil, errors.New("error finding requested resource")
	}

	seriesResponses := mappers.MapSeriesCollectionToSeriesResponseCollection(userSeries)
	return seriesResponses, nil
}

func (raveSeriesService *raveSeriesService) GetSeriesOrganizer(seriesId uint64) (uint64, error) {
	series, err := raveSeriesService.GetById(seriesId)
	if err != nil {
		return 0, errors.New("failed to find series with given id")
	}
	return series.OrganizerID, nil
}

func (raveSeriesService *raveSeriesService) UpdateSeries(seriesId uint64, series *dtos.UpdateSeriesRequest) (*response.SeriesResponse, error) {
	foundSeries, err := raveSeriesService.GetById(seriesId)
	if err != nil {
		return nil, errors.New("failed to find series")
	}
	foundSeries.Name = series.Name
	foundSeries.ImageUrl = series.ImageUrl
	foundSeries.Logo = series.SeriesLogo
	foundSeries.Description = series.Description
	savedSeries, err := raveSeriesService.Save(foundSeries)
	if err != nil {
		return nil, errors.New("failed to update series")
	}
	seriesResponse := &response.SeriesResponse{
		ID:          savedSeries.ID,
		Name:        savedSeries.Name,
		Events:      mappers.MapEventsToEventResponses(savedSeries.Events, savedSeries),
		OrganizerID: savedSeries.OrganizerID,
		ImageUrl:    savedSeries.ImageUrl,
		Description: savedSeries.Description,
		Logo:        savedSeries.Logo,
	}
	return seriesResponse, nil
}

func (raveSeriesService *raveSeriesService) SetEventService(eventService EventService) {
	raveSeriesService.EventService = eventService
}

func (raveSeriesService *raveSeriesService) AddToSeries(seriesId, eventId uint64) (*response.SeriesResponse, error) {
	series, err := raveSeriesService.GetById(seriesId)
	if err != nil {
		return nil, errors.New("failed to find series")
	}
	event, err := raveSeriesService.EventService.GetEventBy(eventId)
	if err != nil {
		return nil, errors.New("failed to find event")
	}
	eventSeries, err := raveSeriesService.GetById(event.SeriesID)
	if err != nil {
		return nil, errors.New("failed to find events series")
	}
	if series.OrganizerID != eventSeries.OrganizerID {
		return nil, errors.New("user not allowed to add event to series")
	}
	event.SeriesID = series.ID
	series.Events = append(series.Events, event)
	err = raveSeriesService.UpdateEvent(event)
	if err != nil {
		return nil, errors.New("failed to add event to series")
	}
	series, err = raveSeriesService.Save(series)
	if err != nil {
		return nil, errors.New("failed to add event to series")
	}
	return mapSeriesToSeriesResponse(series), nil
}
