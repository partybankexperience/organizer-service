package repositories

import (
	"github.com/djfemz/rave/app/models"
)

type OrganizerRepository interface {
	Repository[models.Organizer, uint64]
}

func NewOrganizerRepository() OrganizerRepository {
	var organizerRepository OrganizerRepository = &repositoryImpl[models.Organizer, uint64]{}
	return organizerRepository
}
