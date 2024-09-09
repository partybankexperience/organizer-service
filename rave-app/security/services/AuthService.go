package services

import (
	"bytes"
	"errors"
	request "github.com/djfemz/rave/rave-app/dtos/request"
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/security"
	"github.com/djfemz/rave/rave-app/security/otp"
	"github.com/djfemz/rave/rave-app/services"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
)

type AuthService struct {
	organizerService services.OrganizerService
	attendeeService  services.AttendeeService
	mailService      services.MailService
}

func NewAuthService(organizerService services.OrganizerService,
	attendeeService services.AttendeeService,
	mailService services.MailService) *AuthService {
	return &AuthService{
		organizerService: organizerService,
		attendeeService:  attendeeService,
		mailService:      mailService,
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
		content, err := getMailTemplate(password.Code)
		if err != nil {
			return nil, err
		}
		mailService := services.NewMailService()
		_, err = mailService.Send(request.NewEmailNotificationRequest(org.Username, services.CreateNewOrganizerEmail(content.String())))
		if err != nil {
			return nil, err
		}
		return createAuthResponse(org), nil
	}
}

func (authenticationService *AuthService) ValidateOtp(otp string) (*response.RaveResponse[map[string]any], error) {
	organizerService := authenticationService.organizerService
	org, err := organizerService.GetByOtp(otp)
	if err != nil {
		return nil, err
	}
	orgResponse := mapOrgToOrgResponse(org)
	log.Println("orgResponse: ", orgResponse)
	token, err := security.GenerateAccessTokenFor(org)
	if err != nil {
		return nil, err
	}
	res := map[string]any{
		"token": token,
		"user":  orgResponse,
	}
	return &response.RaveResponse[map[string]any]{Data: res}, nil
}

func (authenticationService *AuthService) AuthenticateAttendee(authRequest request.AttendeeAuthRequest) (*response.LoginResponse, error) {
	attendee, err := authenticationService.attendeeService.GetAttendeeByUsername(authRequest.Username)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(attendee.Password), []byte(authRequest.Password))
	if err != nil {
		return nil, errors.New("authentication failed for attendee")
	}
	return &response.LoginResponse{
		Message:  "user authenticated successfully",
		Username: attendee.Username,
	}, nil
}

func addUser(authRequest *request.AuthRequest, err error, organizerService services.OrganizerService, org *models.Organizer) (*response.LoginResponse, error) {
	organizer, err := organizerService.Create(&authRequest.CreateUserRequest)
	if err != nil {
		log.Println("Error: ", err)
	}
	org, err = organizerService.GetByUsername(organizer.Username)
	if err != nil {
		log.Println("Error: ", err)
	}
	return createAuthResponse(org), nil
}

func createAuthResponse(org *models.Organizer) *response.LoginResponse {
	return &response.LoginResponse{
		Username: org.Username,
		Message:  "check your email for one-time-password",
	}
}

func getMailTemplate(data string) (*bytes.Buffer, error) {
	mailTemplate, err := template.ParseFiles("rave-mail-template.html")
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	err = mailTemplate.Execute(&body, data)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func mapOrgToOrgResponse(organizer *models.Organizer) (orgResponse *response.OrganizationResponse) {
	var series = make([]*response.SeriesResponse, 0)
	orgResponse = &response.OrganizationResponse{}
	orgResponse.UserResponse = &response.UserResponse{}
	orgResponse.ID = organizer.ID
	orgResponse.Username = organizer.Username
	orgResponse.CreatedAt = organizer.CreatedAt
	orgResponse.Name = organizer.Name
	orgResponse.Role = organizer.Role

	log.Println("series: ", organizer.Series)
	for _, orgSeries := range organizer.Series {
		createdSeries := &response.SeriesResponse{
			ID:          orgSeries.ID,
			Name:        orgSeries.Name,
			ImageUrl:    orgSeries.ImageUrl,
			Description: orgSeries.Description,
		}
		series = append(series, createdSeries)
	}
	orgResponse.Series = series
	return orgResponse
}
