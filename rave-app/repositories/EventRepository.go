package repositories

import (
	"github.com/djfemz/rave/rave-app/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	crudRepository[models.Event, uint64]
	FindAllByCalendar(calendarId uint64, pageNumber, pageSize int) ([]*models.Event, error)
}

type raveEventRepository struct {
	*repositoryImpl[models.Event, uint64]
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &raveEventRepository{
		&repositoryImpl[models.Event, uint64]{},
	}
}

func (raveEventRepository *raveEventRepository) FindAllByCalendar(calendarId uint64, pageNumber, pageSize int) ([]*models.Event, error) {
	if pageSize < 1 {
		pageSize = 1
	}
	if pageNumber < 1 {
		pageNumber = 1
	} else if pageSize > 100 {
		pageSize = 100
	}
	offset := (pageNumber - 1) * pageSize
	var events []*models.Event
	db := Connect()
	err := db.Where(&models.Event{SeriesID: calendarId}).Offset(offset).Limit(pageSize).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
