package repositories

import "github.com/djfemz/rave/rave-app/models"

type EventStaffRepository interface {
	crudRepository[models.EventStaff, uint64]
}

type raveEventStaffRepository struct {
	*repositoryImpl[models.EventStaff, uint64]
}

func NewEventStaffRepository() EventStaffRepository {
	return &raveEventStaffRepository{
		&repositoryImpl[models.EventStaff, uint64]{},
	}
}
