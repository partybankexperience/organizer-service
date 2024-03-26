package repositories

import "github.com/djfemz/rave/rave-app/models"

type TicketRepository interface {
	crudRepository[models.Ticket, uint64]
}

type raveTicketRepository struct {
	*repositoryImpl[models.Ticket, uint64]
}

func NewTicketRepository() TicketRepository {
	return &raveTicketRepository{&repositoryImpl[models.Ticket, uint64]{}}
}
