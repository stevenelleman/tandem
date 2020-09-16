package controller

import (
	"sg/services/public-api/models"

	"github.com/gin-gonic/gin"
)

// TODO: Make interface for functions that are used together that require same set of services
// 	interface should correspond to service
type TestController interface {
	ListSilos(ctx *gin.Context) ([]*models.Silo, error)
	GetSilo(*gin.Context, string) (*models.Silo, error)
}

func NewTestController(c *Controller) TestController {
	return c
}
