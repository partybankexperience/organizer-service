package services

import (
	"bytes"
	"encoding/json"
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/mappers"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"github.com/djfemz/rave/rave-app/utils"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
	"net/http"
	"os"
)

type TicketService interface {
	CreateTicketFor(request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error)
	GetById(id uint64) (*response.TicketResponse, error)
	GetTicketById(id uint64) (*models.Ticket, error)
	GetAllTicketsFor(eventId uint64, pageNumber, pageSize int) ([]*models.Ticket, error)
	UpdateTicketSoldOutBy(reference string) (*response.TicketResponse, error)
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

func (raveTicketService *raveTicketService) CreateTicketFor(request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error) {
	event, err := raveTicketService.GetEventBy(request.EventId)
	if err != nil {
		log.Println("event: ", event)
		return nil, errors.New("event not found")
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
	log.Println("perks: ", ticket.TicketPerks)
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
	go sendNewTicketMessageFor(event)
	createTicketResponse.TicketPerks = savedTicket.TicketPerks
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

func sendNewTicketMessageFor(event *models.Event) {
	ticketMessage := buildTicketMessage(event)
	body, err := json.Marshal(ticketMessage)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, os.Getenv("TICKET_SERVICE_URL"), bytes.NewReader(body))
	req.Header.Add("Content-Type", APPLICATION_JSON_VALUE)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
}

func buildTicketMessage(event *models.Event) *request.NewTicketMessage {
	ticketTypes := extractTicketTypesFrom(event.Tickets)
	return &request.NewTicketMessage{
		Types:        ticketTypes,
		Name:         event.Name,
		Reference:    event.Reference,
		Venue:        event.Venue,
		AttendeeTerm: event.AttendeeTerm,
		Date:         event.EventDate,
		TimeFrame:    event.StartTime,
	}
}

func extractTicketTypesFrom(tickets []*models.Ticket) []*request.TicketType {
	ticketTypes := make([]*request.TicketType, 0)
	for _, ticket := range tickets {
		ticketType := &request.TicketType{
			Reference: ticket.Reference,
			Reserved:  ticket.PurchaseLimit,
			MaxSeats:  ticket.NumberAvailable,
			Name:      ticket.Name,
			Price:     ticket.Price,
			Colour:    ticket.Colour,
			Category:  ticket.Category,
			Stock:     ticket.Stock,
			Capacity:  ticket.Capacity,
			Perks:     ticket.TicketPerks,
		}
		if ticket.ActivePeriod != nil {
			ticketType.SalesEndDate = ticket.ActivePeriod.EndDate
			ticketType.SalesEndTime = ticket.ActivePeriod.EndTime
		}
		ticketTypes = append(ticketTypes, ticketType)
	}
	return ticketTypes
}
