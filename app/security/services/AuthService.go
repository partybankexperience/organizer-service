package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/services"
)

type AuthService struct {
	organizerService services.OrganizerService
	mailService      services.MailService
}

func NewAuthService() *AuthService {
	return &AuthService{
		organizerService: services.NewOrganizerService(),
	}
}

func (authenticationService *AuthService) Authenticate(loginRequest *request.LoginRequest) (*response.LoginResponse, error) {
	org, err := authenticationService.organizerService.GetByUsername(loginRequest.Username)
	if err == nil {
		return &response.LoginResponse{
			Username: org.Username,
			Message:  "check your email for one-time-password",
		}, nil
	}
	return nil, err
}
