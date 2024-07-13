package repositories

import "github.com/djfemz/rave/rave-app/models"

type CalendarRepository interface {
	Repository[models.Calendar, uint64]
	FindPublicCalendarFor(organizer uint64) (*models.Calendar, error)
}

type raveCalendarRepository struct {
	repositoryImpl[models.Calendar, uint64]
}

func NewCalendarRepository() CalendarRepository {
	return &raveCalendarRepository{}
}

func (raveCalendarRepository raveCalendarRepository) FindPublicCalendarFor(organizer uint64) (*models.Calendar, error) {
	foundCalendar := &models.Calendar{}
	err := db.Where(&models.Calendar{Name: "Public", OrganizerID: organizer}).Find(foundCalendar).Error
	if err != nil {
		return nil, err
	}
	return foundCalendar, nil
}
