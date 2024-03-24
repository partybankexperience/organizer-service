package middlewares

import (
	"github.com/djfemz/rave/app/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
		authValue := strings.Split(" ", authHeader)
		token := authValue[len(authValue)-1]

	}
}
