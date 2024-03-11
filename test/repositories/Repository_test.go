package test

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepositoryImpl_Save(t *testing.T) {
	var repository repositories.Repository[models.Organizer, uint64] = &repositories.RepositoryImpl[models.Organizer, uint64]{}
	var savedOrg = repository.Save(&models.Organizer{})
	assert.NotNil(t, savedOrg)
}

func TestRepositoryImpl_FindById(t *testing.T) {

}

func TestRepositoryImpl_FindAll(t *testing.T) {

}

func TestRepositoryImpl_DeleteById(t *testing.T) {

}

func TestGetId(t *testing.T) {
	var event = models.Event{
		ID:   24,
		Name: "John",
		Date: time.Now(),
	}
	var id, _ = repositories.GetId(event)
	assert.Equal(t, uint64(24), id)
}
