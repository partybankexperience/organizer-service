package test

import (
	"fmt"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//import (
//	"github.com/djfemz/rave/partybank-app/repositories"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"testing"
//)
//

//
//func TestFindByEventId(t *testing.T) {
//	tickets, err := ticketRepository.FindAllByEventId(2)
//	log.Println("tickets: ", tickets)
//	assert.Nil(t, err)
//	assert.NotNil(t, tickets)
//	assert.NotEmpty(t, tickets)
//}

func TestDeleteTicketsWithEventId(t *testing.T) {
	var ticketRepository = repositories.NewTicketRepository(repositories.Connect())
	var eventId uint64 = 8
	err := ticketRepository.DeleteTicketsFor(eventId)
	assert.Nil(t, err)
}

func TestFindTicketsWithEventId(t *testing.T) {
	var ticketRepository = repositories.NewTicketRepository(repositories.Connect())
	var eventId uint64 = 8
	tickets, _ := ticketRepository.FindAllByEventId(eventId, 1, 30)
	//assert.Nil(t, err)
	log.Println("tickets: ", tickets)
	assert.Empty(t, tickets)
}

func TestDeleteAllNotIn(t *testing.T) {
	db := repositories.Connect()
	db = db.Debug()
	var ticketRepository = repositories.NewTicketRepository(db)
	var eventId uint64 = 8
	tickets, err := ticketRepository.FindAllByEventId(eventId, 1, 30)
	ticketCountBeforeDelete := len(tickets)
	assert.Nil(t, err)
	tickets = tickets[1:]
	err = ticketRepository.DeleteAllNotIn(tickets[0].EventID, tickets)
	//fmt.Println("Error: ", err.Error())
	assert.Nil(t, err)
	tickets, err = ticketRepository.FindAllByEventId(eventId, 1, 30)
	assert.Nil(t, err)
	ticketCountAfterDelete := len(tickets)
	assert.Less(t, ticketCountAfterDelete, ticketCountBeforeDelete)
}

func TestEx(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	nums = nums[1:]
	fmt.Println(nums)
}
