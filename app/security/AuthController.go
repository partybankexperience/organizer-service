package security

import (
	"github.com/djfemz/rave/app/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var organizationService services.OrganizerService = &services.AppOrganizerService{}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(ctx *gin.Context) {
	var loginRequest loginRequest
	var loginResponse = make(map[string]string)
	err := ctx.BindJSON(&loginRequest)
	log.Println(loginRequest)
	if err != nil {
		loginResponse["error"] = err.Error()
		ctx.IndentedJSON(http.StatusBadRequest, loginResponse)
		return
	}
	org := organizationService.GetByUsername(loginRequest.Username)
	if bcrypt.CompareHashAndPassword([]byte(org.Password), []byte(loginRequest.Password)) != nil {
		loginResponse["error"] = err.Error()
		ctx.IndentedJSON(http.StatusBadRequest, loginResponse)
		return
	}
	token, err := GenerateAccessToken(org)
	loginResponse["access_token"] = token
	ctx.JSON(http.StatusOK, loginResponse)
}
