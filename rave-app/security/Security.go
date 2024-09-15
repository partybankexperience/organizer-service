package security

import (
	"errors"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/repositories"

	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strconv"
	"time"
)

type payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

const APP_NAME = "Partybank"

func GenerateAccessTokenFor(user *models.User) (string, error) {
	log.Println("user: ", user)
	token := *jwt.NewWithClaims(jwt.SigningMethodHS256, buildJwtClaimsFor(user))

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

//func GenerateAccessTokenForAttendee(user *models.Attendee) (string, error) {
//	log.Println("user: ", user)
//	token := *jwt.NewWithClaims(jwt.SigningMethodHS256, buildJwtClaimsFor(user.User))
//
//	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
//	if err != nil {
//		return "", err
//	}
//	return accessToken, nil
//}

func ExtractUserFrom(token string) (*models.Organizer, error) {
	db := repositories.Connect()
	organizerRepository := repositories.NewOrganizerRepository(db)
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	}, jwt.WithIssuer(APP_NAME), jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, errors.New("access token is not valid")
	}

	subject, err := tok.Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	org, err := organizerRepository.FindByUsername(subject)
	if err != nil {
		return nil, err
	}
	//claims, err := tok.Claims.GetAudience()
	//if slices.Contains(claims, org.Role) {
	//	return nil, errors.New("user is not authorized to access this resource")
	//}
	return org, nil
}

func buildJwtClaimsFor(user *models.User) *jwt.RegisteredClaims {
	return &jwt.RegisteredClaims{
		Issuer:    APP_NAME,
		Subject:   user.Username,
		Audience:  []string{user.Role, strconv.FormatUint(user.ID, 10)},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
}

func buildJwtClaimsForAttendee(user *models.Attendee) *jwt.RegisteredClaims {
	return &jwt.RegisteredClaims{
		Issuer:    APP_NAME,
		Subject:   user.Username,
		Audience:  []string{user.Role, strconv.FormatUint(user.ID, 10)},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365)),
	}
}
