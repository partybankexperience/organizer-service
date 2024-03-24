package test

import (
	"github.com/djfemz/rave/app/security/otp"
	"github.com/djfemz/rave/app/services"
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
