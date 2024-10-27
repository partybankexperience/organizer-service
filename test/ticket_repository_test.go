package test

import (
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
	tickets, err := ticketRepository.FindAllByEventId(eventId, 1, 30)
	log.Println("tickets: ", tickets)
	assert.Nil(t, err)
	assert.Empty(t, tickets)
}
