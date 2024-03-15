package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
)

type OrganizerService interface {
	Create(createOrganizerRequest *request.CreateOrganizerRequest) *response.CreateOrganizerResponse
}
