package repositories

import "github.com/djfemz/rave/rave-app/models"

type EventRepository interface {
	crudRepository[models.Event, uint64]
	FindAllByCalendar(calendarId uint64) ([]*models.Event, error)
}

type raveEventRepository struct {
	*repositoryImpl[models.Event, uint64]
}

func NewEventRepository() EventRepository {
	return &raveEventRepository{
		&repositoryImpl[models.Event, uint64]{},
	}
}

func (raveEventRepository *raveEventRepository) FindAllByCalendar(calendarId uint64) ([]*models.Event, error) {
	var events []*models.Event
	db := connect()
	err := db.Where(&models.Event{CalendarID: calendarId}).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
