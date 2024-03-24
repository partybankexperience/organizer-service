package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/repositories"
)

type EventStaffService interface {
	Create(createUserRequest *request.CreateUserRequest) (*response.RaveResponse[any], error)
}

type raveEventStaffService struct {
	repositories.EventStaffRepository
	EventService
}

func (eventStaffService *raveEventStaffService) Create(createUserRequest *request.CreateUserRequest) (*response.RaveResponse[any], error) {
	//repo := eventStaffService.EventStaffRepository
	return nil, nil
}

func mapCreateUserToEventStaff(userRequest *request.CreateUserRequest) *models.EventStaff {
	return &models.EventStaff{}
}
