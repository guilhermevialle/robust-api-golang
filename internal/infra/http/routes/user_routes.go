package routes

import (
	controllers "api/internal/infra/http/controllers/customer_controller"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(rg *gin.RouterGroup, customerController controllers.ICustomerController) {
	rg.POST("/customer/register", customerController.CreateCustomer)
}
