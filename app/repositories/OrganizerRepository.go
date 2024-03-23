package repositories

import (
	"github.com/djfemz/rave/app/models"
)

type OrganizerRepository interface {
	CrudRepository[models.Organizer, uint64]
}

func NewOrganizerRepository() OrganizerRepository {
	var organizerRepository OrganizerRepository = &RepositoryImpl[models.Organizer, uint64]{}
	return organizerRepository
}
