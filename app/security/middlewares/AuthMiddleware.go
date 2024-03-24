package middlewares

import (
	response "github.com/djfemz/rave/app/dtos/response"
	"github.com/djfemz/rave/app/security"
	"github.com/djfemz/rave/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
)

var appPaths = make(map[string][]string)

func NewPathToAuthority(path string, authorities ...string) {
	appPaths[path] = authorities
}

func Routers(router *gin.Engine) {
	protected := router.Group("/protected")
	{
		protected.POST("")
		protected.POST("")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
		authValue := strings.Split(" ", authHeader)
		token := authValue[len(authValue)-1]
		org, err := security.ExtractUserFrom(token)
		if err != nil {
			ctx.JSON(http.StatusForbidden, &response.RaveResponse[string]{Data: "token is invalid"})
			return
		}
		path := ctx.FullPath()
		if org != nil && slices.Contains(appPaths[path], org.Role) {
			ctx.Next()
		}
	}
}
