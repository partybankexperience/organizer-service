package test

import (
	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/djfemz/rave/rave-app/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUpdateAttendee(t *testing.T) {
	username := utils.GenerateTicketReference()
	attendee := &models.Attendee{
		FullName: "John Doe",
		User: &models.User{
			Username: username,
			Role:     models.ATTENDEE,
		},
		PhoneNumber: "09087654356",
	}
	db := repositories.Connect()
	attendeeRepository := repositories.NewAttendeeRepository(db)
	attendeeService := services.NewAttendeeService(attendeeRepository, nil)
	attendee, _ = attendeeRepository.Save(attendee)
	updateAttendeeRequest := &dtos.UpdateAttendeeRequest{
		FullName:    "James Doe",
		PhoneNumber: "09087654356",
	}

	res, err := attendeeService.UpdateAttendee(attendee.ID, updateAttendeeRequest)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	attendee, err = attendeeRepository.FindById(attendee.ID)
	log.Println("attendee: ", *attendee)
	assert.Nil(t, err)
	assert.NotNil(t, attendee)
	assert.Equal(t, username, attendee.Username)
	assert.Equal(t, "James Doe", attendee.FullName)
}
