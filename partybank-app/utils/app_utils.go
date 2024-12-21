package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/google/uuid"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	AUTHORIZATION           = "Authorization"
	APP_NAME                = "Partybank"
	APP_EMAIL               = "partybankexperience@gmail.com"
	FRONT_END_TEST_BASE_URL = "https://partybank-dev.vercel.app"
	FRONT_END_DEV_BASE_URL  = "http://localhost:5173"
	FRONT_END_PROD_URL      = "https://thepartybank.com"
	EVENT_REFERENCE_PREFIX  = "evt-"
	TICKET_REFERENCE_PREFIX = "tkt-"
	GET_IMAGE_URL_PATH      = "/api/v1/image/"
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

func SendNewTicketMessageFor(event *models.Event) {
	ticketMessage := BuildTicketMessage(event)
	body, err := json.Marshal(ticketMessage)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("request body: ", string(body))
	req, err := http.NewRequest(http.MethodPost, os.Getenv("ADD_EVENT_ENDPOINT_PAYMENT_SERVICE"), bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	log.Println("request data: ", req.Body)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("response: ", *res)
}

func BuildTicketMessage(event *models.Event) *request.NewTicketMessage {
	ticketTypes := extractTicketTypesFrom(event.Tickets)
	var timeFrame = event.StartTime + " to " + event.EndTime

	return &request.NewTicketMessage{
		Types:                 ticketTypes,
		Name:                  event.Name,
		Reference:             event.Reference,
		Venue:                 event.Venue,
		IsNotificationEnabled: event.IsNotificationEnabled,
		PhoneNumber:           event.ContactInformation,
		AttendeeTerm:          event.AttendeeTerm,
		Date:                  event.EventDate,
		TimeFrame:             timeFrame,
		CreatedBy:             event.CreatedBy,
	}
}

func extractTicketTypesFrom(tickets []*models.Ticket) []*request.TicketType {
	ticketTypes := make([]*request.TicketType, 0)
	for _, ticket := range tickets {
		ticketType := &request.TicketType{
			Reference:           ticket.Reference,
			Name:                ticket.Name,
			Price:               strconv.FormatFloat(ticket.Price, 'f', -1, 64),
			Color:               ticket.Colour,
			Category:            strings.ToLower(ticket.Category),
			GroupTicketCapacity: ticket.GroupTicketCapacity,
			Stock:               ToTitleCase(ToTitleCase(ticket.Stock)),
			Capacity:            int(ticket.Capacity),
			Perks:               strings.Join(ticket.TicketPerks, ","),
			Type:                ToTitleCase(ticket.Type),
			PurchaseLimit:       int(ticket.PurchaseLimit),
		}
		if ticket.ActivePeriod != nil {
			ticketType.SalesEndDate = ticket.ActivePeriod.EndDate
			ticketType.SalesEndTime = ticket.ActivePeriod.EndTime
		}
		ticketTypes = append(ticketTypes, ticketType)
	}
	return ticketTypes
}

func ToTitleCase(text string) string {
	return cases.Title(language.English).String(text)
}

func GenerateEventReference() string {
	uniqueId := uuid.New()
	uniqueHash := base64.RawURLEncoding.EncodeToString([]byte(uniqueId.String()))
	return EVENT_REFERENCE_PREFIX + uniqueHash
}

func GenerateTicketReference() string {
	uniqueId := uuid.New()
	uniqueHash := base64.RawURLEncoding.EncodeToString([]byte(uniqueId.String()))
	return TICKET_REFERENCE_PREFIX + uniqueHash
}

func GenerateImageId(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func ExistsWithTicketName(event *models.Event, name string) bool {
	for _, ticket := range event.Tickets {
		if strings.EqualFold(ticket.Name, name) {
			return true
		}
	}
	return false
}
