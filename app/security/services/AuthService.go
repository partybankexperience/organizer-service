package services

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/security/otp"
	"github.com/djfemz/rave/app/services"
	"log"
)

type AuthService struct {
	organizerService services.OrganizerService
	mailService      services.MailService
}

func NewAuthService() *AuthService {
	return &AuthService{
		organizerService: services.NewOrganizerService(),
		mailService:      services.NewMailService(),
	}
}

func (authenticationService *AuthService) Authenticate(authRequest *request.AuthRequest) (*response.LoginResponse, error) {
	organizerService := authenticationService.organizerService
	org, err := organizerService.GetByUsername(authRequest.Username)
	if err != nil {
		res, err := addUser(authRequest, err, organizerService, org)
		return res, err
	} else {
		password := otp.GenerateOtp()
		_, err = organizerService.UpdateOtpFor(org.ID, password)
		if err != nil {
			return nil, err
		}
		authenticationService.mailService.Send(request.NewEmailNotificationRequest(org.Username, services.CreateNewOrganizerEmail(password.Code)))
		return createAuthResponse(org), nil
	}
}

func addUser(authRequest *request.AuthRequest, err error, organizerService services.OrganizerService, org *models.Organizer) (*response.LoginResponse, error) {
	organizer, err := organizerService.Create(&authRequest.CreateOrganizerRequest)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	org, err = organizerService.GetByUsername(organizer.Username)
	if err != nil {
		log.Fatal("Error: ", err)

	}
	return createAuthResponse(org), err
}

func createAuthResponse(org *models.Organizer) *response.LoginResponse {
	return &response.LoginResponse{
		Username: org.Username,
		Message:  "check your email for one-time-password",
	}
}
