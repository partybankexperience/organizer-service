package security

import (
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var organizationService services.OrganizerService = &services.AppOrganizerService{}
var err error

type loginRequest struct {
	Username string `json:"username"`
}

func LoginHandler(ctx *gin.Context) {
	var loginRequest loginRequest

	if err = ctx.BindJSON(&loginRequest); err != nil {
		handleError(ctx, err)
		return
	}
	org, err := organizationService.GetByUsername(loginRequest.Username)

	token, err := GenerateAccessToken(org)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, &response.RaveResponse[string]{Data: token})
}

func handleError(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusBadRequest, &response.RaveResponse[error]{Data: err})
}
