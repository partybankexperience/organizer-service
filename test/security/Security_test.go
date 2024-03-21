package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/security"
	"github.com/djfemz/rave/app/security/otp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var organization = &models.Organizer{
	User: &models.User{
		Username: "jon@email.com",
		Role:     models.ORGANIZER,
	},
	CreatedAt: time.Now(),
}

func TestGenerateToken(t *testing.T) {
	accessToken, err := security.GenerateAccessToken(organization)
	assert.NotNil(t, accessToken)
	assert.Nil(t, err)
	assert.NotEmpty(t, accessToken)
}

func TestValidateOtp(t *testing.T) {
	otp := otp.GenerateOtp("jon@email.com")
	security.ValidateOtp("")
}
