package repositories

import "github.com/djfemz/rave/rave-app/models"

type TicketRepository interface {
	crudRepository[models.Ticket, uint64]
	FindAllByEventId(id uint64, pageNumber, pageSize int) ([]*models.Ticket, error)
}

type raveTicketRepository struct {
	*repositoryImpl[models.Ticket, uint64]
}

func NewTicketRepository() TicketRepository {
	return &raveTicketRepository{&repositoryImpl[models.Ticket, uint64]{}}
}

func (raveTicketRepository *raveTicketRepository) FindAllByEventId(id uint64, pageNumber, pageSize int) ([]*models.Ticket, error) {
	if pageSize < 1 {
		pageSize = 1
	}
	if pageNumber < 1 {
		pageNumber = 1
	} else if pageSize > 100 {
		pageSize = 100
	}
	offset := (pageNumber - 1) * pageSize
	var tickets []*models.Ticket
	db := connect()
	err := db.Where(&models.Ticket{EventId: id}).Offset(offset).Limit(pageSize).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
