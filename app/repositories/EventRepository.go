package repositories

import "github.com/djfemz/rave/app/models"

type EventRepository interface {
	crudRepository[models.Event, uint64]
}

type raveEventRepository struct {
	*repositoryImpl[models.Event, uint64]
}

func NewEventRepository() EventRepository {
	return &repositoryImpl[models.Event, uint64]{}
}
