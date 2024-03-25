package controllers

import (
	request "github.com/djfemz/rave/app/dtos/request"
	response "github.com/djfemz/rave/app/dtos/response"
	authService "github.com/djfemz/rave/app/security/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var err error

type AuthController struct {
	*authService.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService.NewAuthService(),
	}
}

func (authController *AuthController) AuthHandler(ctx *gin.Context) {
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

func (authController *AuthController) ValidateOtp(ctx *gin.Context) {
	log.Println("ctx: ", ctx)
	code := ctx.Query("code")
	log.Println("code: ", code)

	res, err := authController.AuthService.ValidateOtp(code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.LoginResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func handleError(ctx *gin.Context, err error) {
	ctx.IndentedJSON(http.StatusBadRequest, &response.RaveResponse[string]{Data: err.Error()})
}
