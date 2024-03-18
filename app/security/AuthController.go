package security

import (
	"github.com/djfemz/rave/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var organizationService services.OrganizerService = &services.AppOrganizerService{}
var loginResponse = make(map[string]string)
var err error

type loginRequest struct {
	Username string `json:"username"`
}

func LoginHandler(ctx *gin.Context) {
	var loginRequest loginRequest

	if err = ctx.BindJSON(&loginRequest); err != nil {
		handleError(ctx, err, loginResponse)
		return
	}
	org := organizationService.GetByUsername(loginRequest.Username)

	token, err := GenerateAccessToken(org)
	if err != nil {
		handleError(ctx, err, loginResponse)
		return
	}
	loginResponse["access_token"] = token
	ctx.JSON(http.StatusOK, loginResponse)
}

func handleError(ctx *gin.Context, err error, loginResponse map[string]string) {
	loginResponse["error"] = err.Error()
	ctx.IndentedJSON(http.StatusBadRequest, loginResponse)
}
