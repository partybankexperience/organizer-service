package test

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var eventService = services.NewEventService()

func TestCreateEvent(t *testing.T) {
	createEvent := &request.CreateEventRequest{
		Name:               "test event",
		Location:           "Sabo Yaba",
		Date:               "2024-03-23",
		Time:               "12:00:00",
		ContactInformation: "09023456789",
		Description:        "this is a test event",
	}

	res, err := eventService.Create(createEvent)
	log.Println("event response: ", *res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestGetEventById(t *testing.T) {
	event, err := eventService.GetById(1)
	assert.NotNil(t, event)
	assert.Nil(t, err)
}

func TestEditEventDetails(t *testing.T) {

}
