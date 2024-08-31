package repositories

import (
	"github.com/djfemz/rave/rave-app/models"
	"gorm.io/gorm"
)

type EventStaffRepository interface {
	crudRepository[models.EventStaff, uint64]
}

type raveEventStaffRepository struct {
	*repositoryImpl[models.EventStaff, uint64]
}

func NewEventStaffRepository(db *gorm.DB) EventStaffRepository {
	return &raveEventStaffRepository{
		&repositoryImpl[models.EventStaff, uint64]{
			db,
		},
	}
}
