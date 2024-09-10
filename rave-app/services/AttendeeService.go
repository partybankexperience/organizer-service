package services

import (
	"bytes"
	"errors"
	"github.com/djfemz/rave/rave-app/security"
	"html/template"

	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/mappers"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"log"
)

type AttendeeService interface {
	Register(createAttendeeRequest *dtos.CreateAttendeeRequest) (*response.AttendeeResponse, error)
	GetAttendeeByUsername(username string) (*response.AttendeeResponse, error)
}

type raveAttendeeService struct {
	repositories.AttendeeRepository
	MailService
}

func NewAttendeeService(attendeeRepository repositories.AttendeeRepository, mailService MailService) AttendeeService {
	return &raveAttendeeService{
		AttendeeRepository: attendeeRepository,
		MailService:        mailService,
	}
}

func (attendeeService *raveAttendeeService) Register(createAttendeeRequest *dtos.CreateAttendeeRequest) (*response.AttendeeResponse, error) {
	attendee := mappers.MapCreateAttendeeRequestToAttendee(createAttendeeRequest)
	attendee, err := attendeeService.Save(attendee)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, errors.New("failed to create attendee service")
	}
	attendeeWelcomeMailRequest, err := buildNewAttendeeMessageFor(attendee)
	if err != nil {
		return nil, err
	}
	go attendeeService.MailService.Send(attendeeWelcomeMailRequest)

	return mappers.MapAttendeeToAttendeeResponse(attendee), nil

}

func (attendeeService *raveAttendeeService) GetAttendeeByUsername(username string) (*response.AttendeeResponse, error) {
	attendee, err := attendeeService.FindByUsername(username)
	if err != nil {
		return nil, errors.New("failed to find attendee")
	}
	return mappers.MapAttendeeToAttendeeResponse(attendee), nil
}

func buildNewAttendeeMessageFor(attendee *models.Attendee) (*dtos.EmailNotificationRequest, error) {
	templ, err := getAttendeeEmailTemplate(attendee)
	if err != nil {
		return nil, errors.New("could not get mail template")
	}
	return &dtos.EmailNotificationRequest{
		Sender: dtos.Sender{
			Name:  "Partybank",
			Email: "partybankexperience@gmail.com",
		},
		Recipients: []dtos.Recipient{
			{
				Name:  attendee.FullName,
				Email: attendee.Username,
			},
		},
		Subject: "Welcome mail",
		Content: templ,
	}, nil
}

type attendeeMessage struct {
	FullName string
	Link     string
}

func getAttendeeEmailTemplate(attendee *models.Attendee) (string, error) {
	token, err := security.GenerateAccessTokenForAttendee(attendee)
	if err != nil {
		return "", err
	}
	message := &attendeeMessage{
		FullName: attendee.FullName,
		Link:     "https://www.thepartybank.com?" + "token=" + token,
	}
	mailTemplate, err := template.ParseFiles("rave-mail-template-new.html")
	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}
	var body bytes.Buffer
	err = mailTemplate.Execute(&body, message)
	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}
	return body.String(), nil
}
