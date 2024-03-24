package test

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
	"github.com/djfemz/rave/app/security/otp"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var repository = repositories.NewOrganizerRepository()

func TestRepositoryImpl_Save(t *testing.T) {
	username := "johnny@email.com"
	var savedOrg = repository.Save(&models.Organizer{
		Name:      "John",
		CreatedAt: time.Now(),
		User: &models.User{
			Username: username,
		},
		Otp: otp.GenerateOtp(),
	})
	log.Println(savedOrg)
	assert.NotNil(t, savedOrg)
}

func TestFindByUsername(t *testing.T) {
	found, _ := repository.FindByUsername("johnny@email.com")
	log.Println(found)
	assert.NotNil(t, found)
}

func TestRepositoryImpl_FindById(t *testing.T) {
	foundOrg := repository.FindById(3)
	log.Println(foundOrg)
	assert.NotNil(t, foundOrg)
}

func TestRepositoryImpl_FindAll(t *testing.T) {
	orgs := repository.FindAll()
	log.Println(orgs)
	assert.Equal(t, 3, len(orgs))

}

func TestRepository_FindAll_Pagination(t *testing.T) {
	var pageable = repositories.NewPageAble(1, 1)
	orgs := repository.FindAllBy(pageable)
	assert.NotNil(t, orgs)
	assert.Equal(t, 1, len(orgs))
}

func TestRepositoryImpl_DeleteById(t *testing.T) {
	repository.DeleteById(1)
	orgs := repository.FindAll()
	assert.Equal(t, 2, len(orgs))
}

func TestFindByOtp(t *testing.T) {
	org, err := repository.FindByOtp("506554")
	assert.NotNil(t, org)
	assert.Nil(t, err)
}

func TestGetId(t *testing.T) {
	var event = models.Event{
		ID:   24,
		Name: "John",
		Date: time.Now().String(),
	}
	var id, _ = repositories.GetId(event)
	assert.Equal(t, uint64(24), id)
}
