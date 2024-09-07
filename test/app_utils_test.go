package test

import (
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/utils"
	"log"
	"testing"
)

func TestIsValidEndTime(t *testing.T) {
	ticket := &models.Ticket{
		SaleEndDate:  "2024-03-01",
		SalesEndTime: "09:00:00",
	}
	status := utils.IsTicketSaleEndedFor(ticket)
	log.Println(status)
}
