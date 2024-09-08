package utils

import (
	"encoding/base64"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

const (
	AUTHORIZATION = "Authorization"
)

func ConvertQueryStringToInt(query string) (int, error) {
	value, err := strconv.Atoi(query)
	if err != nil {
		log.Println("Error: ", err)
		return 0, err
	}
	return value, nil
}

func isDateValid(date string) bool {
	return false
}

func GenerateEventReference() string {
	s := uuid.New()
	v := base64.RawURLEncoding.EncodeToString([]byte(s.String()))
	return "evt-" + v
}

func GenerateTicketReference() string {
	s := uuid.New()
	v := base64.RawURLEncoding.EncodeToString([]byte(s.String()))
	return "tkt-" + v
}

func IsTicketSaleEndedFor(ticket *models.Ticket) bool {
	if ticket.ActivePeriod == nil {
		return false
	}
	ticketEndTime := ticket.ActivePeriod.EndDate + " " + ticket.ActivePeriod.EndTime
	endTime, err := time.Parse("2006-01-02 15:04:05", ticketEndTime)
	if err != nil {
		log.Println("err: ", err)
		return false
	}
	log.Println("true: ", endTime)
	return time.Now().After(endTime)
}

func ExistsWithTicketName(event *models.Event, name string) bool {
	for _, ticket := range event.Tickets {
		if ticket.Name == name {
			return true
		}
	}
	return false
}
