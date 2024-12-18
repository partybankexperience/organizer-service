package integrations

import (
	"errors"
	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/utils"
	"log"
	"net/http"
	"os"
	"time"
)

type PaymentService interface {
	Authenticate() (token string, err error)
	IsAccessTokenInvalid() bool
	DeleteEventBy(reference string) error
}

type PartyBankPaymentService struct {
	clientId     string
	clientSecret string
	AccessToken  *response.PaymentServiceAccessToken
}

func NewPaymentService(clientId, clientSecret string) PaymentService {
	return &PartyBankPaymentService{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
func (partyBankPaymentService *PartyBankPaymentService) Authenticate() (token string, err error) {
	client := utils.NewHttpClient[request.PaymentServiceAuthRequest, response.PaymentServiceAuthResponse](&log.Logger{}, &http.Client{})
	paymentServiceAuthRequest := &request.PaymentServiceAuthRequest{
		ClientId:     partyBankPaymentService.clientId,
		ClientSecret: partyBankPaymentService.clientSecret,
	}
	res, err := client.Send(http.MethodPost, os.Getenv("PAYMENT_SERVICE_AUTH_URL"), paymentServiceAuthRequest, nil)
	if err != nil {
		return "", err
	}
	log.Println("res:", *res)
	if res.Data.AccessToken == "" {
		return "", errors.New("failed to authenticate with payment service, try again later")
	}
	partyBankPaymentService.AccessToken = &response.PaymentServiceAccessToken{
		AccessToken: res.Data.AccessToken,
		ExpiresAt:   time.Now().AddDate(0, 0, 15),
	}
	return res.Data.AccessToken, nil
}

func (partyBankPaymentService *PartyBankPaymentService) IsAccessTokenInvalid() bool {
	isAccessTokenInvalid := partyBankPaymentService.AccessToken == nil ||
		partyBankPaymentService.AccessToken.ExpiresAt.Before(time.Now())
	if isAccessTokenInvalid {
		return true
	}
	return false
}

func (partyBankPaymentService *PartyBankPaymentService) DeleteEventBy(reference string) error {
	paymentServiceDeleteEndpoint := os.Getenv("PAYMENT_SERVICE_DELETE_EVENT_URL")
	paymentServiceDeleteEndpoint = paymentServiceDeleteEndpoint + reference
	if partyBankPaymentService.IsAccessTokenInvalid() {
		_, err := partyBankPaymentService.Authenticate()
		if err != nil {
			log.Println("error obtaining token from payment service")
			return errors.New("error obtaining token from payment service")
		}
	}
	client := utils.NewHttpClient[interface{}, interface{}](&log.Logger{}, &http.Client{})
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + partyBankPaymentService.AccessToken.AccessToken
	_, err := client.Send(http.MethodPost, paymentServiceDeleteEndpoint, nil, headers)
	if err != nil {
		log.Println("Error deleting event on payment service: ", err)
		return errors.New("error deleting event on payment service")
	}
	return nil
}
