package di

import (
	app_services "api/internal/app/services"
	"api/internal/app/use_cases"
	"api/internal/infra/http/controllers"
	"api/internal/infra/repositories"
	infra_services "api/internal/infra/services"
)

type Container struct {
	CustomerController controllers.ICustomerController
	AuthController     controllers.IAuthController
}

func NewContainer() *Container {
	customerRepo := repositories.NewCustomerRepo()
	tokenService := infra_services.NewTokenService()
	customerService := use_cases.NewCustomerService(customerRepo)
	authService := app_services.NewAuthService(customerService, tokenService)

	return &Container{
		CustomerController: controllers.NewCustomerController(customerService),
		AuthController:     controllers.NewAuthController(authService),
	}
}
