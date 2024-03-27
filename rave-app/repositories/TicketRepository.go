package repositories

import "github.com/djfemz/rave/rave-app/models"

type TicketRepository interface {
	crudRepository[models.Ticket, uint64]
	FindAllByEventId(id uint64) ([]*models.Ticket, error)
}

type raveTicketRepository struct {
	*repositoryImpl[models.Ticket, uint64]
}

func NewTicketRepository() TicketRepository {
	return &raveTicketRepository{&repositoryImpl[models.Ticket, uint64]{}}
}

func (raveTicketRepository *raveTicketRepository) FindAllByEventId(id uint64) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	db := connect()
	err := db.Where(&models.Ticket{EventId: id}).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
