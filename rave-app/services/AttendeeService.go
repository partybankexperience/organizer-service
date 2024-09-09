package services

import (
	"errors"
	"fmt"

	"log"
	"os"

	dtos "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/mappers"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/utils"
)

type AttendeeService interface {
	Register(createAttendeeRequest *dtos.CreateAttendeeRequest) (*response.AttendeeResponse, error)
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

func (attendService *raveAttendeeService) Register(createAttendeeRequest *dtos.CreateAttendeeRequest) (*response.AttendeeResponse, error) {
	attendee := mappers.MapCreateAttendeeRequestToAttendee(createAttendeeRequest)
	attendee, err := attendService.Save(attendee)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, errors.New("failed to create attendee service")
	}
	attendeeWelcomeMailRequest := buildNewAttendeeMessageFor(attendee)
	go attendService.MailService.Send(attendeeWelcomeMailRequest)

	return mappers.MapAttendeeToAttendeeResponse(attendee), nil

}

func buildNewAttendeeMessageFor(attendee *models.Attendee) *dtos.EmailNotificationRequest {
	return &dtos.EmailNotificationRequest{
		Sender: dtos.Sender{
			Name:  utils.APP_NAME,
			Email: utils.APP_EMAIL,
		},
		Recipients: []dtos.Recipient{
			{
				Name:  attendee.FullName,
				Email: attendee.Username,
			},
		},
		Subject: "Welcome mail",
		Content: getAttendeeEmailTemplate(attendee),
	}
}

func getAttendeeEmailTemplate(attendee *models.Attendee) string {
	mail, err := os.ReadFile("rave-mail-template-new.html")
	if err != nil {
		log.Fatal("Error: ", err)
		return ""
	}
	return fmt.Sprintf(string(mail), attendee.FullName)
}
