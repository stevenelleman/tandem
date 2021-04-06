package controller

import (
	"web-microservice-shell/libraries/golang/guts/models"

	"github.com/gin-gonic/gin"
)

func (c *BaseController) ListSilos(ctx *gin.Context) ([]*models.Silo, error) {
	err := c.SayHello(ctx)
	if err != nil {
		return nil, err
	}

	return c.psqlConn.ListSilos(ctx)
}

func (c *BaseController) GetSilo(ctx *gin.Context, id string) (*models.Silo, error) {
	return c.psqlConn.GetSilo(ctx, id)
}

func (c *BaseController) CreateSilo(ctx *gin.Context, s *models.Silo) error {
	return c.psqlConn.CreateSilo(ctx, s)
}

func (c *BaseController) UpdateSilo(ctx *gin.Context, s *models.Silo) error {
	return c.psqlConn.UpdateSilo(ctx, s)
}

func (c *BaseController) DeleteSilo(ctx *gin.Context, id string) error {
	return c.psqlConn.DeleteSilo(ctx, id)
}

// Dummy method to demonstrate that grpc connection to sg service works
func (c *BaseController) SayHello(ctx *gin.Context) error {
	return c.sgConn.SayHello(ctx)
}
