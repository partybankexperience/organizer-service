package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/djfemz/rave/app/services"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func GenerateAccessTokenFor(user *models.Organizer) (string, error) {
	token := *jwt.NewWithClaims(jwt.SigningMethodHS256, buildJwtClaimsFor(user))
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func ExtractUserFrom(token string) (*models.Organizer, error) {
	var organizerService = services.NewOrganizerService()
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	subject, err := tok.Claims.GetSubject()
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	org, err := organizerService.GetByUsername(subject)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	return org, nil
}

func buildJwtClaimsFor(user *models.Organizer) *jwt.RegisteredClaims {
	return &jwt.RegisteredClaims{
		Issuer:    "app",
		Subject:   user.Username,
		Audience:  []string{user.Role},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
}
