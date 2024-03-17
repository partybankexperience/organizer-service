package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/security"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var user = &models.Organizer{
	Username:  "jon@email.com",
	Password:  "password",
	Role:      models.ORGANIZER,
	CreatedAt: time.Now(),
}

func TestName(t *testing.T) {
	access_token, err := security.GenerateAccessToken(user)
	log.Println(access_token)
	assert.NotNil(t, access_token)
	assert.Nil(t, err)
	assert.NotEmpty(t, access_token)

}
