package controllers

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	authService "github.com/djfemz/rave/app/security/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error

type AutController struct {
	*authService.AuthService
}

func NewAuthController() *AutController {
	return &AutController{
		authService.NewAuthService(),
	}
}

func (authController *AutController) LoginHandler(ctx *gin.Context) {
	var signInRequest request.AuthRequest
	if err = ctx.BindJSON(&signInRequest); err != nil {
		handleError(ctx, err)
		return
	}

	res, err := authController.AuthService.Authenticate(&signInRequest)
	if res != nil {
		ctx.JSON(http.StatusOK, response.RaveResponse[response.LoginResponse]{Data: *res})
	} else {
		handleError(ctx, err)
	}
}

func handleError(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
}
