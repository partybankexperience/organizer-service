package services

import (
	"errors"
	"fmt"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"log"
)

type DiscountService interface {
	CreateDiscount(request *request.CreateDiscountRequest) (*response.CreateDiscountResponse, error)
}

type raveDiscountService struct {
	repositories.DiscountRepository
	TicketService
}

func NewDiscountService() DiscountService {
	return &raveDiscountService{
		repositories.NewDiscountRepository(),
		NewTicketService(),
	}
}

func (raveDiscountService *raveDiscountService) CreateDiscount(request *request.CreateDiscountRequest) (*response.CreateDiscountResponse, error) {
	ticketId := request.TicketId
	ticket, err := ticketRepository.FindById(ticketId)
	if err != nil || ticket == nil {
		return nil, errors.New(fmt.Sprintf("Cannot create Discount\n reason: Failed to find ticket with id: %d", ticketId))
	}
	discount := &models.Discount{}
	createDiscountResponse := &response.CreateDiscountResponse{}
	model.Copy(discount, request)
	discount.Ticket = ticket
	savedDiscount, err := raveDiscountService.Save(discount)
	log.Println("saved discount: ", savedDiscount)
	if err != nil {
		return nil, err
	}
	model.Copy(createDiscountResponse, savedDiscount)
	return createDiscountResponse, nil
}
