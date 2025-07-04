package app

import (
	"api/internal/infra/di"
	"api/internal/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	r := gin.Default()
	c := di.NewContainer()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.CustomerRoutes(r.Group("/"), c.CustomerController)
	routes.AuthRoutes(r.Group("/"), c.AuthController)

	return r
}
