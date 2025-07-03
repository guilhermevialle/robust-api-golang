package di

import (
	"api/internal/app/use_cases"
	controllers "api/internal/infra/http/controllers/customer_controller"
	"api/internal/infra/repositories"
)

type Container struct {
	CustomerController controllers.ICustomerController
}

func NewContainer() *Container {
	return &Container{
		CustomerController: setupCustomerDI(),
	}
}

func setupCustomerDI() controllers.ICustomerController {
	customerRepo := repositories.NewCustomerRepo()
	customerService := use_cases.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	return customerController
}
