package repositories

import (
	"errors"
	"github.com/djfemz/organizer-service/partybank-app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type EventRepository interface {
	crudRepository[models.Event, uint64]
	FindEventById(id uint64) (*models.Event, error)
	FindAllByCalendar(calendarId uint64, pageNumber, pageSize int) ([]*models.Event, error)
	FindAllPublishedByPage(page int, size int) ([]*models.Event, error)
	FindByReference(reference string) (*models.Event, error)
	FindAllByOrganizer(organizerId uint64, pageNumber, pageSize int) ([]*models.Event, error)
	DeleteEventById(eventId uint64) error
}

type raveEventRepository struct {
	*repositoryImpl[models.Event, uint64]
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &raveEventRepository{
		&repositoryImpl[models.Event, uint64]{
			db,
		},
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
	err := raveEventRepository.Db.Where(&models.Event{SeriesID: calendarId, IsEventDeleted: false}).Offset(offset).Limit(pageSize).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (raveEventRepository *raveEventRepository) FindEventById(id uint64) (*models.Event, error) {
	event := models.Event{}
	err := raveEventRepository.Db.Preload(clause.Associations).
		Where("id=? AND is_event_deleted=?", id, false).First(&event).Error
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, errors.New("event with given id not found")
	}
	return &event, err
}

func (raveEventRepository *raveEventRepository) FindAllPublishedByPage(page int, size int) ([]*models.Event, error) {
	if size < 1 {
		size = 1
	}
	if page < 1 {
		page = 1
	} else if size > 100 {
		size = 100
	}
	offset := (page - 1) * size
	log.Println("offset: ", offset)
	var events []*models.Event
	err := raveEventRepository.Db.Preload(clause.Associations).
		Where(&models.Event{PublicationState: models.PUBLISHED, IsEventDeleted: false}).
		Offset(offset).Limit(size).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (raveEventRepository *raveEventRepository) FindByReference(reference string) (*models.Event, error) {
	event := &models.Event{}
	err := raveEventRepository.Db.Preload(clause.Associations).Where(&models.Event{Reference: reference, IsEventDeleted: false}).First(event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (raveEventRepository *raveEventRepository) FindAllByOrganizer(organizerId uint64, page, size int) ([]*models.Event, error) {
	if size < 1 {
		size = 1
	}
	if page < 1 {
		page = 1
	} else if size > 100 {
		size = 100
	}
	offset := (page - 1) * size
	log.Println("offset: ", offset)
	var events []*models.Event

	err := raveEventRepository.Db.Preload(clause.Associations).
		Joins("JOIN series ON series.id = events.series_id").
		Where("series.organizer_id = ?", organizerId).
		Offset(offset).
		Limit(size).
		Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (raveEventRepository *raveEventRepository) DeleteEventById(eventId uint64) error {
	event := &models.Event{ID: eventId}
	err := raveEventRepository.Db.Delete(event).Error
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil
}
