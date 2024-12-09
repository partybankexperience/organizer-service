package services

import (
	"crypto/tls"
	"errors"
	request "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

type goMailService struct {
}

func NewGoMailService() MailService {
	return &goMailService{}
}

func (goMailService *goMailService) Send(emailRequest *request.EmailNotificationRequest) (string, error) {
	message := gomail.NewMessage()
	message.SetHeader("From", emailRequest.Sender.Email)
	message.SetHeader("To", emailRequest.Recipients[0].Email, emailRequest.Recipients[0].Email)
	message.SetAddressHeader("Cc", emailRequest.Recipients[0].Email, emailRequest.Recipients[0].Email)
	message.SetHeader("Subject", emailRequest.Subject)
	message.SetBody("text/html", emailRequest.Content)
	//message.Attach("/home/Alex/lolcat.jpg")

	//dialer := gomail.NewDialer("smtp.gmail.com", 587, "partybankexperience@gmail.com", "mlztvopuabontioo")
	dialer := gomail.NewDialer("smtp.gmail.com", 465,
		os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		log.Println("Error: ", err)
		return "", errors.New("error sending email")
	}
	return "mail sent successfully", nil
}
