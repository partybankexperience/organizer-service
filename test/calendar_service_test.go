package test

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var calendarService services.SeriesService

func TestCreateCalendar(t *testing.T) {
	calendarService = services.NewSeriesService()
	createCalendarRequest := &request.CreateCalendarRequest{
		Name:        "test",
		OrganizerID: 1,
		Description: "test desc",
		ImageUrl:    "https://image.com",
	}

	response, err := calendarService.CreateCalendar(createCalendarRequest)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Message)
	assert.Nil(t, err)
}

func TestGetCalendar(t *testing.T) {
	calendarService = services.NewSeriesService()
	calendar, err := calendarService.GetById(1)
	log.Println("calendar: ", calendar.Events)
	assert.Nil(t, err)
	assert.NotNil(t, calendar)
}

func TestAddEventToCalendar(t *testing.T) {
	req := &request.CreateEventRequest{
		CalendarId: 1,
		Name:       "rave",
		Location:   "Abuja",
		Time:       time.Now().String(),
	}
	eventService = services.NewEventService()
	event, err := eventService.Create(req)
	assert.Nil(t, err)
	calendarService = services.NewSeriesService()
	calendar, err := calendarService.AddEventToCalendar(1, event)
	assert.Nil(t, err)
	assert.NotNil(t, calendar)
	assert.NotEmpty(t, calendar.Events)
	assert.GreaterOrEqual(t, len(calendar.Events), 1)

}
