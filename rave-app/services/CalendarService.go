package services

import (
	"errors"
	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
)

type CalendarService interface {
	CreateCalendar(createCalendarRequest *dtos.CreateCalendarRequest) (*response.CreateCalendarResponse, error)
	GetById(id uint64) (*models.Calendar, error)
	AddEventToCalendar(id uint64, event *models.Event) (*models.Calendar, error)
}

type raveCalendarService struct {
	repositories.CalendarRepository
}

func NewCalendarService() CalendarService {
	return &raveCalendarService{repositories.NewCalendarRepository()}
}

func (raveCalendarService *raveCalendarService) CreateCalendar(createCalendarRequest *dtos.CreateCalendarRequest) (*response.CreateCalendarResponse, error) {
	calendar := &models.Calendar{}
	errs := model.Copy(calendar, createCalendarRequest)
	isCopyErrorPresent := len(errs) > 0
	if isCopyErrorPresent {
		return nil, errors.New("error creating calendar from request")
	}
	savedCalendar, err := raveCalendarService.CalendarRepository.Save(calendar)
	if err != nil {
		return nil, err
	}
	createdCalendarResponse := &response.CreateCalendarResponse{}
	createdCalendarResponse.ID = savedCalendar.ID
	createdCalendarResponse.Name = savedCalendar.Name
	createdCalendarResponse.Message = "calendar created successfully"
	return createdCalendarResponse, nil
}

func (raveCalendarService *raveCalendarService) GetById(id uint64) (*models.Calendar, error) {
	calendar, err := raveCalendarService.CalendarRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (raveCalendarService *raveCalendarService) AddEventToCalendar(id uint64, event *models.Event) (*models.Calendar, error) {
	calendar, err := raveCalendarService.GetById(id)
	if err != nil {
		return nil, err
	}
	calendar.Events = append(calendar.Events, event)
	calendar, err = raveCalendarService.CalendarRepository.Save(calendar)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}
