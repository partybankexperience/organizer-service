package services

import (
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type TicketService interface {
	CreateTicketFor(request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error)
	GetTicketById(id uint64) (*response.TicketResponse, error)
	GetAllTicketsFor(eventId uint64) ([]*models.Ticket, error)
}

var ticketRepository = repositories.NewTicketRepository()

type raveTicketService struct {
}

func NewTicketService() TicketService {
	return &raveTicketService{}
}

func (raveTicketService *raveTicketService) CreateTicketFor(request *request.CreateTicketRequest) (addTicketResponse *response.TicketResponse, err error) {
	eventService := NewEventService()
	event, err := eventService.GetEventBy(request.EventId)

	ticket := &models.Ticket{}
	errs := model.Copy(ticket, request)
	if len(errs) != 0 {
		log.Println(errs)
		return nil, errors.New("failed to create ticket")
	}
	savedTicket, err := ticketRepository.Save(ticket)
	if err != nil {
		log.Println("error: ", err)
		return nil, errors.New("failed to save ticket")
	}
	event.Tickets = append(event.Tickets, ticket)
	log.Println("ticket: ", savedTicket)
	err = eventService.UpdateEvent(event)
	if err != nil {
		return nil, errors.New("failed to save ticket")
	}
	createTicketResponse := &response.TicketResponse{}
	errs = model.Copy(createTicketResponse, savedTicket)
	return createTicketResponse, nil
}

func (raveTicketService *raveTicketService) GetTicketById(id uint64) (*response.TicketResponse, error) {
	ticket, err := ticketRepository.FindById(id)
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

func (raveTicketService *raveTicketService) GetAllTicketsFor(eventId uint64) ([]*models.Ticket, error) {
	tickets, err := ticketRepository.FindAllByEventId(eventId)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
