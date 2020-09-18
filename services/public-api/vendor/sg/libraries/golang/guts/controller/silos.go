package controller

import (
	"sg/libraries/golang/guts/handlers/requests"
	"sg/libraries/golang/guts/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PsqlClose() {
	c.psqlConn.Close()
}

func (c *Controller) SgClose() {
	c.sgConn.Close()
}

func (c *Controller) ListSilos(ctx *gin.Context) ([]*models.Silo, error) {
	err := c.sgConn.SayHello(ctx)
	if err != nil {
		return nil, err
	}

	return c.psqlConn.ListSilos(ctx)
}

func (c *Controller) GetSilo(ctx *gin.Context, id string) (*models.Silo, error) {
	return c.psqlConn.GetSilo(ctx, id)
}

func (c *Controller) CreateSilo(ctx *gin.Context, s *requests.Silo) error {
	return c.psqlConn.CreateSilo(ctx, s)
}

func (c *Controller) UpdateSilo(ctx *gin.Context, s *requests.Silo) error {
	return c.psqlConn.UpdateSilo(ctx, s)
}

func (c *Controller) DeleteSilo(ctx *gin.Context, id string) error {
	return c.psqlConn.DeleteSilo(ctx, id)
}
