package utils

import (
	"encoding/base64"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
)

const (
	AUTHORIZATION           = "Authorization"
	APP_NAME                = "Partybank"
	APP_EMAIL               = "partybankexperience@gmail.com"
	FRONT_END_TEST_BASE_URL = "https://partybank-dev.vercel.app"
	FRONT_END_DEV_BASE_URL  = "http://localhost:5173"
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

func ExistsWithTicketName(event *models.Event, name string) bool {
	for _, ticket := range event.Tickets {
		if strings.EqualFold(ticket.Name, name) {
			return true
		}
	}
	return false
}
