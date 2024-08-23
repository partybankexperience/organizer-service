package repositories

import "github.com/djfemz/rave/rave-app/models"

type CalendarRepository interface {
	Repository[models.Series, uint64]
	FindPublicCalendarFor(organizer uint64) (*models.Series, error)
}

type raveCalendarRepository struct {
	repositoryImpl[models.Series, uint64]
}

func NewCalendarRepository() CalendarRepository {
	return &raveCalendarRepository{}
}

func (raveCalendarRepository raveCalendarRepository) FindPublicCalendarFor(organizer uint64) (*models.Series, error) {
	foundCalendar := &models.Series{}
	err := db.Where(&models.Series{Name: "Public", OrganizerID: organizer}).Find(foundCalendar).Error
	if err != nil {
		return nil, err
	}
	return foundCalendar, nil
}
