package test

import (
	request "github.com/djfemz/rave/rave-app/dtos/request"
	"github.com/djfemz/rave/rave-app/security/otp"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var organizerService = services.NewOrganizerService()

func TestUpdateOtpForOrganizer(t *testing.T) {
	testOtp := otp.GenerateOtp()
	organizer, _ := organizerService.UpdateOtpFor(14, testOtp)
	assert.NotNil(t, organizer)
	assert.Equal(t, organizer.Otp, testOtp)
}

func TestGetById(t *testing.T) {
	organizer, err := organizerService.GetById(23)
	assert.NotNil(t, organizer)
	assert.Nil(t, err)
}

func TestOrganizerCanAddEventStaff(t *testing.T) {
	addEventStaff := &request.AddEventStaffRequest{
		StaffEmails: []string{"test@email.com"},
	}
	response, err := organizerService.AddEventStaff(addEventStaff)
	assert.NotNil(t, response)
	assert.Nil(t, err)
}

func TestOrganizerCanAddEvent(t *testing.T) {
	addEventRequest := &request.CreateEventRequest{
		Name:               "test event",
		Location:           "Sabo Yaba",
		Date:               "2024-03-23",
		Time:               "12:00:00",
		ContactInformation: "09023456789",
		Description:        "this is a test event",
		OrganizerId:        23,
	}
	response, err := organizerService.AddEvent(addEventRequest)
	assert.NotNil(t, response)
	assert.Nil(t, err)
}
