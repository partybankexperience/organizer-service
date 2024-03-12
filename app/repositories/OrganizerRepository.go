package repositories

import (
	"github.com/djfemz/rave/app/models"
)

type OrganizerRepository interface {
	Repository[models.Organizer, uint64]
}

type OrganizerRepositoryImpl struct {
}

func NewOrganizerRepositoryImpl() OrganizerRepository {
	var rep OrganizerRepository = &RepositoryImpl[models.Organizer, uint64]{}
	return rep
}
