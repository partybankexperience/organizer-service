package test

import (
	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	"github.com/djfemz/organizer-service/partybank-app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mailService = services.NewGoMailService()

func TestSendMail(t *testing.T) {
	var mailRequest = buildMailRequest()
	response, _ := mailService.Send(mailRequest)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response)
}

func buildMailRequest() *request.EmailNotificationRequest {
	return &request.EmailNotificationRequest{
		Sender: request.Sender{
			Email: "partybankexperience@gmail.com",
			Name:  "partybank",
		},
		Recipients: []request.Recipient{
			{
				Email: "oladejifemi00@gmail.com",
				Name:  "oladejifemi00@gmail.com",
			},
		},
		Subject: "Hello",
		Content: "<p>Hello Friend, how are you doing?</p>",
	}
}
