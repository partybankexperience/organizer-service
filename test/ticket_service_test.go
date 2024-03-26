package test

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var addTicketRequest = &request.CreateTicketRequest{
	Type:            "single",
	Name:            "early birds",
	NumberAvailable: 50,
	Price:           5000.00,
	EventId:         2,
}

var ticketService = services.NewTicketService()

func TestAddTicket(t *testing.T) {
	res, err := ticketService.CreateTicketFor(addTicketRequest)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
