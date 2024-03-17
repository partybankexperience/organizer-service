package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/security"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var user = &models.Organizer{
		Username:  "jon@email.com",
		Password:  "password",
		Role:      models.ORGANIZER,
		CreatedAt: time.Now(),
	}
	assert.NotEmpty(t, security.GenerateAccessToken(user))
}
