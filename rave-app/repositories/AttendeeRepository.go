package repositories

import (
	"github.com/djfemz/rave/rave-app/models"
	"gorm.io/gorm"
)

type AttendeeRepository interface {
	crudRepository[models.Attendee, uint64]
}

type raveAttendeeRepository struct {
	*repositoryImpl[models.Attendee, uint64]
}

func NewAttendeeRepository(db *gorm.DB) AttendeeRepository {
	return &raveAttendeeRepository{
		&repositoryImpl[models.Attendee, uint64]{
			db,
		},
	}
}
