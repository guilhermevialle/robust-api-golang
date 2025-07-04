package middlewares

import (
	env_config "api/internal/infra/config"
	infra_services "api/internal/infra/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var env *env_config.EnvVariables = env_config.LoadEnv()

type IAuthMiddleware interface {
	Handle() gin.HandlerFunc
}

type AuthMiddleware struct {
	tokenService infra_services.ITokenService
}

var _ IAuthMiddleware = (*AuthMiddleware)(nil)

func NewAuthMiddleware(tokenService infra_services.ITokenService) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: tokenService,
	}
}

func (am *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		isValid, err := am.tokenService.Validate(tokenString, env.TokenSecret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
