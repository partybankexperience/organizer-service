package services

import (
	"errors"
	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type SeriesService interface {
	CreateCalendar(createCalendarRequest *dtos.CreateCalendarRequest) (*response.CreateCalendarResponse, error)
	GetById(id uint64) (*models.Series, error)
	AddEventToCalendar(id uint64, event *models.Event) (*models.Series, error)
	GetCalendar(id uint64) (*response.CreateCalendarResponse, error)
	GetPublicCalendarFor(id uint64) (*models.Series, error)
}

type raveSeriesService struct {
	repositories.SeriesRepository
}

func NewSeriesService() SeriesService {
	return &raveSeriesService{repositories.NewSeriesRepository()}
}

func (raveCalendarService *raveSeriesService) CreateCalendar(createCalendarRequest *dtos.CreateCalendarRequest) (*response.CreateCalendarResponse, error) {
	calendar := &models.Series{}
	errs := model.Copy(calendar, createCalendarRequest)
	log.Println("calendar: ", *calendar)
	isCopyErrorPresent := len(errs) > 0
	if isCopyErrorPresent {
		return nil, errors.New("error creating calendar from request")
	}
	savedCalendar, err := raveCalendarService.SeriesRepository.Save(calendar)
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

func (raveCalendarService *raveSeriesService) GetById(id uint64) (*models.Series, error) {
	calendar, err := raveCalendarService.SeriesRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (raveCalendarService *raveSeriesService) GetCalendar(id uint64) (*response.CreateCalendarResponse, error) {
	resp := &response.CreateCalendarResponse{}
	calendar, err := raveCalendarService.GetById(id)
	if err != nil {
		return nil, err
	}
	errs := model.Copy(resp, calendar)
	if len(errs) > 0 {
		return nil, err
	}
	return resp, nil
}

func (raveCalendarService *raveSeriesService) AddEventToCalendar(id uint64, event *models.Event) (*models.Series, error) {
	calendar, err := raveCalendarService.GetById(id)
	if err != nil {
		return nil, err
	}
	log.Println("event: ", *event)
	calendar.Events = append(calendar.Events, event)
	log.Println("calendar: ", calendar.Events)
	calendar, err = raveCalendarService.SeriesRepository.Save(calendar)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (raveCalendarService *raveSeriesService) GetPublicCalendarFor(id uint64) (*models.Series, error) {
	calendar, err := raveCalendarService.SeriesRepository.FindPublicCalendarFor(id)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}
