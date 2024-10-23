package services

import (
	"bytes"
	"encoding/json"
	"errors"
	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/mappers"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
	"github.com/djfemz/organizer-service/partybank-app/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type TicketService interface {
	CreateTicketFor(eventId uint64, request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error)
	AddTickets(eventId uint64, tickets []*request.CreateTicketRequest) ([]*response.TicketResponse, error)
	GetById(id uint64) (*response.TicketResponse, error)
	GetTicketById(id uint64) (*models.Ticket, error)
	GetAllTicketsFor(eventId uint64, pageNumber, pageSize int) ([]*models.Ticket, error)
	UpdateTicketSoldOutBy(reference string) (*response.TicketResponse, error)
	UpdateTicket(ticketId uint64, updateTicketRequest *request.UpdateTicketRequest) (*response.TicketResponse, error)
	EditTicket(ticketId uint64, editTicketRequest *request.EditTicketRequest) (editTicketResponse *response.TicketResponse, err error)
}

type raveTicketService struct {
	repositories.TicketRepository
	EventService
}

func NewTicketService(ticketRepository repositories.TicketRepository, eventService EventService) TicketService {
	return &raveTicketService{
		ticketRepository,
		eventService,
	}
}

func (raveTicketService *raveTicketService) CreateTicketFor(eventId uint64, request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error) {
	event, err := raveTicketService.GetEventBy(eventId)
	if err != nil {
		log.Println("event: ", event)
		return nil, errors.New("event not found")
	}
	if utils.ExistsWithTicketName(event, request.Name) {
		return nil, errors.New("ticket already exists")
	}
	ticket := &models.Ticket{}
	errs := model.Copy(ticket, request)
	if len(errs) != 0 {
		log.Println(errs)
		return nil, errors.New("failed to create ticket")
	}
	ticket.Reference = utils.GenerateTicketReference()
	ticket.EventID = event.ID
	ticket.TicketPerks = request.TicketPerks
	ticket.ActivePeriod = &models.ActivePeriod{
		StartDate: request.SalesStartDate,
		EndDate:   request.SaleEndDate,
		StartTime: request.SalesStartTime,
		EndTime:   request.SalesEndTime,
	}
	savedTicket, err := raveTicketService.TicketRepository.Save(ticket)
	if err != nil {
		log.Println("error: ticket saving failed", err)
		return nil, errors.New("failed to save ticket")
	}
	event.Tickets = append(event.Tickets, savedTicket)
	err = raveTicketService.UpdateEvent(event)
	if err != nil {
		return nil, errors.New("failed to save ticket")
	}
	createTicketResponse := &response.TicketResponse{}
	errs = model.Copy(createTicketResponse, savedTicket)
	log.Println("new ticket created: ", savedTicket.TicketPerks)
	//go sendNewTicketMessageFor(event)
	createTicketResponse.TicketPerks = savedTicket.TicketPerks
	if savedTicket.ActivePeriod != nil {
		createTicketResponse.SalesStartDate = savedTicket.ActivePeriod.StartDate
		createTicketResponse.SalesStartTime = savedTicket.ActivePeriod.StartTime
		createTicketResponse.SalesEndTime = savedTicket.ActivePeriod.EndTime
		createTicketResponse.SaleEndDate = savedTicket.ActivePeriod.EndDate
	}
	return createTicketResponse, nil
}

func (raveTicketService *raveTicketService) GetById(id uint64) (*response.TicketResponse, error) {
	ticket, err := raveTicketService.FindById(id)
	if err != nil {
		return nil, err
	}
	res := &response.TicketResponse{}
	errs := model.Copy(res, ticket)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	res.TicketPerks = ticket.TicketPerks
	return res, nil
}

func (raveTicketService *raveTicketService) GetTicketById(id uint64) (*models.Ticket, error) {
	ticket, err := raveTicketService.FindById(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (raveTicketService *raveTicketService) UpdateTicketSoldOutBy(reference string) (*response.TicketResponse, error) {
	ticket, err := raveTicketService.TicketRepository.FindTicketByReference(reference)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	ticket.IsSoldOutTicket = true
	ticket, err = raveTicketService.TicketRepository.Save(ticket)
	if err != nil {
		log.Println("Error: saving failed: ", err)
		return nil, errors.New(err.Error())
	}
	return mappers.MapTicketToTicketResponse(ticket), nil
}

func (raveTicketService *raveTicketService) GetAllTicketsFor(eventId uint64, pageNumber, pageSize int) ([]*models.Ticket, error) {
	tickets, err := raveTicketService.FindAllByEventId(eventId, pageNumber, pageSize)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return tickets, nil
}

func (raveTicketService *raveTicketService) AddTickets(eventId uint64, tickets []*request.CreateTicketRequest) ([]*response.TicketResponse, error) {
	res := make([]*response.TicketResponse, 0)
	for _, ticketRequest := range tickets {
		ticketResponse, _ := raveTicketService.CreateTicketFor(eventId, ticketRequest)
		res = append(res, ticketResponse)
	}
	event, err := raveTicketService.EventService.GetEventBy(eventId)
	if err != nil {
		log.Println("error: ticket_service.go: 146 \t\tfailed to find event")
	}
	go sendNewTicketMessageFor(event)
	return res, nil
}

func (raveTicketService *raveTicketService) UpdateTicket(ticketId uint64, updateTicketRequest *request.UpdateTicketRequest) (*response.TicketResponse, error) {
	foundTicket, err := raveTicketService.GetTicketById(ticketId)
	if err != nil {
		return nil, errors.New("failed to find ticket")
	}
	foundTicket.Name = updateTicketRequest.Name
	foundTicket.Type = updateTicketRequest.Type
	foundTicket.PurchaseLimit = updateTicketRequest.PurchaseLimit
	foundTicket.Capacity = updateTicketRequest.Capacity
	foundTicket.Colour = updateTicketRequest.Colour
	foundTicket.Price = updateTicketRequest.Price
	foundTicket.Stock = updateTicketRequest.Stock
	foundTicket.IsTransferPaymentFeesToGuest = updateTicketRequest.IsTransferPaymentFeesToGuest
	foundTicket.Category = updateTicketRequest.Category
	if foundTicket.ActivePeriod != nil {
		foundTicket.ActivePeriod.StartTime = updateTicketRequest.SalesStartTime
		foundTicket.ActivePeriod.EndTime = updateTicketRequest.SalesEndTime
		foundTicket.ActivePeriod.StartDate = updateTicketRequest.SalesStartDate
		foundTicket.ActivePeriod.EndDate = updateTicketRequest.SaleEndDate
	}
	foundTicket, err = raveTicketService.Save(foundTicket)
	if err != nil {
		return nil, errors.New("failed to save ticket")
	}
	event, err := raveTicketService.EventService.GetEventBy(foundTicket.EventID)
	go sendNewTicketMessageFor(event)
	return mappers.MapTicketToTicketResponse(foundTicket), nil
}

func (raveTicketService *raveTicketService) EditTicket(ticketId uint64, editTicketRequest *request.EditTicketRequest) (editTicketResponse *response.TicketResponse, err error) {
	ticket, err := raveTicketService.GetTicketById(ticketId)
	if err != nil {
		return nil, errors.New("failed to find ticket")
	}
	ticket = mappers.MapEditTicketRequestToTicket(editTicketRequest, ticket)
	ticketResponse := mappers.MapTicketToTicketResponse(ticket)
	return ticketResponse, nil
}

func sendNewTicketMessageFor(event *models.Event) {
	ticketMessage := buildTicketMessage(event)
	body, err := json.Marshal(ticketMessage)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("request body: ", string(body))
	req, err := http.NewRequest(http.MethodPost, os.Getenv("TICKET_SERVICE_URL"), bytes.NewReader(body))
	req.Header.Add("Content-Type", APPLICATION_JSON_VALUE)
	log.Println("request data: ", req.Body)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("response: ", *res)
}

func buildTicketMessage(event *models.Event) *request.NewTicketMessage {
	ticketTypes := extractTicketTypesFrom(event.Tickets)
	var timeFrame = event.StartTime + " to " + event.EndTime

	return &request.NewTicketMessage{
		Types:        ticketTypes,
		Name:         event.Name,
		Reference:    event.Reference,
		Venue:        event.Venue,
		AttendeeTerm: event.AttendeeTerm,
		Date:         event.EventDate,
		TimeFrame:    timeFrame,
		CreatedBy:    event.CreatedBy,
	}
}

func extractTicketTypesFrom(tickets []*models.Ticket) []*request.TicketType {
	ticketTypes := make([]*request.TicketType, 0)
	for _, ticket := range tickets {
		ticketType := &request.TicketType{
			Reference:     ticket.Reference,
			Name:          ticket.Name,
			Price:         strconv.FormatFloat(ticket.Price, 'f', -1, 64),
			Color:         ticket.Colour,
			Category:      strconv.FormatUint(ticket.Category, 10),
			Stock:         ToTitleCase(ToTitleCase(ticket.Stock)),
			Capacity:      int(ticket.Capacity),
			Perks:         strings.Join(ticket.TicketPerks, ","),
			Type:          ToTitleCase(ticket.Type),
			PurchaseLimit: int(ticket.PurchaseLimit),
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
