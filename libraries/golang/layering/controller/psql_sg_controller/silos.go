package psql_sg_controller

import (
	"sg/libraries/golang/layering/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ListSilos(ctx *gin.Context) ([]*models.Silo, error) {
	err := c.SayHello(ctx)
	if err != nil {
		return nil, err
	}

	return c.psqlConn.ListSilos(ctx)
}

func (c *Controller) GetSilo(ctx *gin.Context, id string) (*models.Silo, error) {
	return c.psqlConn.GetSilo(ctx, id)
}

func (c *Controller) CreateSilo(ctx *gin.Context, s *models.Silo) error {
	return c.psqlConn.CreateSilo(ctx, s)
}

func (c *Controller) UpdateSilo(ctx *gin.Context, s *models.Silo) error {
	return c.psqlConn.UpdateSilo(ctx, s)
}

func (c *Controller) DeleteSilo(ctx *gin.Context, id string) error {
	return c.psqlConn.DeleteSilo(ctx, id)
}

// Dummy method to demonstrate that grpc connection to sg service works
func (c *Controller) SayHello(ctx *gin.Context) error {
	return c.sgConn.SayHello(ctx)
}
