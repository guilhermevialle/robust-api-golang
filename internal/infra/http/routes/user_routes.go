package routes

import (
	"api/internal/infra/http/controllers"
	"api/internal/infra/http/middlewares"
	infra_services "api/internal/infra/services"

	"github.com/gin-gonic/gin"
)

var tokenService = infra_services.NewTokenService()
var am = middlewares.NewAuthMiddleware(tokenService)

func CustomerRoutes(rg *gin.RouterGroup, customerController controllers.ICustomerController) {
	rg.GET("/me/:id", am.Handle(), customerController.GetProfile)
	rg.POST("/customer/register", customerController.CreateCustomer)
}
