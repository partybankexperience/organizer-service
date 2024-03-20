package services

import (
	"bytes"
	"encoding/json"
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/config"
	"log"
	"net/http"
)

const (
	API_KEY_VALUE          = "api-key"
	CONTENT_TYPE_KEY       = "Content-Type"
	APPLICATION_JSON_VALUE = "application/json"
	ACCEPT_HEADER_KEY      = "accept"
)

type MailService interface {
	Send(emailRequest *request.EmailNotificationRequest) string
}

type RaveMailService struct {
}

func (raveMailService *RaveMailService) Send(emailRequest *request.EmailNotificationRequest) string {
	jsonData, _ := json.Marshal(emailRequest)
	appConfig := config.LoadConfigFile()
	req, err := http.NewRequest(http.MethodPost, appConfig.MAIL_API_URL, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal("Error Creating request", err)
	}
	addHeadersTo(req, appConfig)

	client := &http.Client{}
	if _, err = client.Do(req); err != nil {
		log.Fatal("Error sending mail: ", err)
	}
	return response.MAIL_SENDING_SUCCESS_MESSAGE
}

func addHeadersTo(req *http.Request, appConfig *config.EnvConfig) {
	req.Header.Add(API_KEY_VALUE, appConfig.MAIL_API_KEY)
	req.Header.Add(CONTENT_TYPE_KEY, APPLICATION_JSON_VALUE)
	req.Header.Add(ACCEPT_HEADER_KEY, APPLICATION_JSON_VALUE)
}
