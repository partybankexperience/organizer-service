package security

import (
	"errors"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/services"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

type payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

const APP_NAME = "rave"

func GenerateAccessTokenFor(user *models.Organizer) (string, error) {
	token := *jwt.NewWithClaims(jwt.SigningMethodHS256, buildJwtClaimsFor(user))

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func ExtractUserFrom(token string) (*models.Organizer, error) {
	var organizerService = services.NewOrganizerService()
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	}, jwt.WithIssuer(APP_NAME), jwt.WithExpirationRequired())

	if !tok.Valid {
		return nil, errors.New("access token is not valid")
	}
	if err != nil {
		return nil, err
	}
	subject, err := tok.Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	org, err := organizerService.GetByUsername(subject)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func buildJwtClaimsFor(user *models.Organizer) *jwt.RegisteredClaims {
	return &jwt.RegisteredClaims{
		Issuer:    APP_NAME,
		Subject:   user.Username,
		Audience:  []string{user.Role, strconv.FormatUint(user.ID, 10)},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
}
