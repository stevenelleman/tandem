package controller

import (
	"sg/libraries/golang/guts/connection/service/psql_conn"
	"sg/libraries/golang/guts/connection/service/sg_conn"
	"sg/libraries/golang/guts/models"

	"github.com/gin-gonic/gin"
)

type ControllerFactory struct {
	psqlConnFactory *psql_conn.PsqlConnectionFactory
	sgConnFactory   *sg_conn.SgConnectionFactory
}

func NewControllerFactory(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) *ControllerFactory {
	cf := &ControllerFactory{}
	cf.psqlConnFactory = psql_conn.NewPsqlConnFactory(psqlArgs)
	cf.sgConnFactory = sg_conn.NewSgConnFactory(sgArgs)
	return cf
}

// TODO: Need to return Controller interface which uses psql and sg only
func (f *ControllerFactory) Controller() *Controller {
	return &Controller{
		psqlConn: f.psqlConnFactory.Connection(),
		sgConn:   f.sgConnFactory.Connection(),
	}
}

// TODO: Split controller by datastore intersections
// Need connection-level methods though, in theory could make a controller-level method if you wanted
type SgController interface {
	SgClose()
}

type PsqlController interface {
	PsqlClose()

	ListSilos(ctx *gin.Context) ([]*models.Silo, error)
	GetSilo(*gin.Context, string) (*models.Silo, error)
}

// HERE: Only methods available to corresponding services/stores
// HERE: Should connection-level methods be represented as controller-level methods?
type PublicAPIController interface {
	SgController
	PsqlController
}

func (f *ControllerFactory) TestController() PublicAPIController {
	// HERE: Only include necessary connection types for composed controller interface
	return &Controller{
		psqlConn: f.psqlConnFactory.Connection(),
		sgConn:   f.sgConnFactory.Connection(),
	}
}
