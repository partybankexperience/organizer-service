package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
)

type EventService interface {
	Create(createEventRequest *request.CreateEventRequest) (response.RaveResponse[any], error)
}

type RaveEventService struct {
	EventService
}
