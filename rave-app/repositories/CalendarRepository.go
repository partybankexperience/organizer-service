package repositories

import "github.com/djfemz/rave/rave-app/models"

type CalendarRepository interface {
	Repository[models.Calendar, uint64]
}

func NewCalendarRepository() CalendarRepository {
	return &repositoryImpl[models.Calendar, uint64]{}
}
