package security

import (
	"github.com/djfemz/rave/app/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateAccessToken(user *models.Organizer) (string, error) {
	token := *jwt.NewWithClaims(jwt.SigningMethodHS256, buildJwtClaimsFor(user))
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func buildJwtClaimsFor(user *models.Organizer) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		Issuer:    "app",
		Subject:   "access_token",
		Audience:  []string{user.Role},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
}
