package services

import (
	"bytes"
	"encoding/json"
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
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

	ticket := &models.Ticket{}
	errs := model.Copy(ticket, request)
	if len(errs) != 0 {
		log.Println(errs)
		return nil, errors.New("failed to create ticket")
	}
	savedTicket, err := raveTicketService.Save(ticket)
	if err != nil {
		return nil, errors.New("failed to save ticket")
	}
	event.Tickets = append(event.Tickets, ticket)
	event.PublicationState = models.PUBLISHED
	err = raveTicketService.UpdateEvent(event)
	if err != nil {
		return nil, errors.New("failed to save ticket")
	}
	createTicketResponse := &response.TicketResponse{}
	errs = model.Copy(createTicketResponse, savedTicket)
	log.Println("new ticket created: ", savedTicket)
	go sendNewTicketEvent(event, createTicketResponse)
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
	return res, nil
}

func (raveTicketService *raveTicketService) GetTicketById(id uint64) (*models.Ticket, error) {
	ticket, err := raveTicketService.FindById(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (raveTicketService *raveTicketService) GetAllTicketsFor(eventId uint64, pageNumber, pageSize int) ([]*models.Ticket, error) {
	tickets, err := raveTicketService.FindAllByEventId(eventId, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func sendNewTicketEvent(event *models.Event, ticketResponse *response.TicketResponse) {
	ticketMessage := buildTicketMessage(event, ticketResponse)
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
		log.Fatal("Error: ", err)
	}

}

func buildTicketMessage(event *models.Event, ticketResponse *response.TicketResponse) *request.NewTicketMessage {
	return &request.NewTicketMessage{
		Type:                       ticketResponse.Type,
		Name:                       ticketResponse.Name,
		Stock:                      ticketResponse.Stock,
		NumberAvailable:            ticketResponse.NumberAvailable,
		Price:                      ticketResponse.Price,
		DiscountCode:               ticketResponse.DiscountCode,
		DiscountPrice:              ticketResponse.DiscountAmount,
		PurchaseLimit:              ticketResponse.PurchaseLimit,
		Percentage:                 ticketResponse.Percentage,
		AvailableDiscountedTickets: ticketResponse.AvailableDiscountedTickets,
		EventName:                  event.Name,
		Description:                event.Description,
		Location:                   event.Location,
	}

}
