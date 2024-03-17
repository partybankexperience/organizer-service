package security

import (
	"errors"
	"github.com/djfemz/rave/app/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var organizationService services.OrganizerService = &services.AppOrganizerService{}

type loginRequest struct {
	username string
	password string
}

func LoginHandler(ctx *gin.Context) {
	var loginRequest loginRequest
	var loginResponse = make(map[string]string)
	err := ctx.BindJSON(&loginRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.New("invalid request object"))
	}
	org := organizationService.GetByUsername(loginRequest.username)
	if bcrypt.CompareHashAndPassword([]byte(org.Password), []byte(loginRequest.password)) != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.New("invalid authentication credentials"))
	}
	token, err := GenerateAccessToken(org)
	loginResponse["access_token"] = token
	ctx.JSON(http.StatusOK, loginResponse)
}
