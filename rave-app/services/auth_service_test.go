package services

import (
	"github.com/djfemz/rave/rave-app/models"
	"log"
	"testing"
)

func TestBuildMessageForAttendee(t *testing.T) {
	attendee := &models.Attendee{
		FullName: "John Doe",
		User: &models.User{
			Username: "john@email.com",
		},
	}
	req := buildNewAttendeeMessageFor(attendee)
	log.Println("content: ", req.Content)
}
