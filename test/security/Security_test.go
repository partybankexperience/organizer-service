package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/security"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var user = &models.Organizer{
	Username:  "jon@email.com",
	Role:      models.ORGANIZER,
	CreatedAt: time.Now(),
}

func TestGenerateToken(t *testing.T) {
	accessToken, err := security.GenerateAccessToken(user)
	assert.NotNil(t, accessToken)
	assert.Nil(t, err)
	assert.NotEmpty(t, accessToken)
}
