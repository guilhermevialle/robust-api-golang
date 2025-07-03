package controllers

import (
	"api/internal/app/use_cases"
	"api/internal/dtos"

	"github.com/gin-gonic/gin"
)

type ICustomerController interface {
	CreateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
}

type CustomerController struct {
	customerService use_cases.ICustomerService
}

var _ ICustomerController = (*CustomerController)(nil)

func NewCustomerController(customerService use_cases.ICustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (uc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var body dtos.CreateCustomerDto

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"BAD_REQUEST": "Invalid input"})
		return
	}

	customer, err := uc.customerService.CreateCustomer(body.Name, body.Email)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, customer)
}

func (uc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := uc.customerService.DeleteCustomer(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}
