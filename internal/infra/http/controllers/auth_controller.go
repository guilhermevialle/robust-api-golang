package controllers

import (
	app_services "api/internal/app/services"
	"api/internal/dtos"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
}

type AuthController struct {
	authService app_services.IAuthService
}

var _ IAuthController = (*AuthController)(nil)

func NewAuthController(authService app_services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body dtos.LoginDto

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"BAD_REQUEST": "Invalid input fields"})
		return
	}

	tokens, err := c.authService.Login(body.Email, body.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"UNAUTHORIZED": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": tokens[0], "refresh_token": tokens[1]})
}
