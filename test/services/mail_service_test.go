package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	"github.com/djfemz/rave/app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mailService = services.NewMailService()

func TestSendMail(t *testing.T) {
	var mailRequest = buildMailRequest()
	var response = mailService.Send(mailRequest)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response)
}

func buildMailRequest() *request.EmailNotificationRequest {
	return &request.EmailNotificationRequest{
		Sender: request.Sender{
			Email: "john@gmail.com",
			Name:  "rave",
		},
		Recipient: []request.Recipient{
			{
				Email: "davipe1322@irnini.com",
				Name:  "John Doe",
			},
		},
		Subject: "Hello",
		Content: "<p>Hello Friend, how are you doing?</p>",
	}
}
