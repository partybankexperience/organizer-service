package test

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var calendarService services.CalendarService

func TestCreateCalendar(t *testing.T) {
	calendarService = services.NewCalendarService()
	createCalendarRequest := &request.CreateCalendarRequest{
		Name: "test",
	}

	response, err := calendarService.CreateCalendar(createCalendarRequest)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Message)
	assert.Nil(t, err)
}

func TestGetCalendar(t *testing.T) {
	calendarService = services.NewCalendarService()
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
	calendarService = services.NewCalendarService()
	calendar, err := calendarService.AddEventToCalendar(1, event)
	assert.Nil(t, err)
	assert.NotNil(t, calendar)
	assert.NotEmpty(t, calendar.Events)
	assert.Equal(t, 1, len(calendar.Events))

}
