package routes

import (
	"api/internal/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, authController controllers.IAuthController) {
	rg.POST("/auth/login", authController.Login)
}
