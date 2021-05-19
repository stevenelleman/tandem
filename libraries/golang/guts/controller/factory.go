package controller

import (
	"github.com/shell/libraries/golang/guts/connection/service/psql_conn"
	"github.com/shell/libraries/golang/guts/connection/service/sg_conn"
)

type ControllerFactory struct {
	psqlConnFactory *psql_conn.PsqlConnectionFactory
	sgConnFactory   *sg_conn.SgConnectionFactory
}

// I like the idea of passing in and instantiating all connections but empty args will simply return empty object
// that way the Controller() for BaseController will still work, but the connection may not be live
func NewControllerFactory(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) *ControllerFactory {
	cf := &ControllerFactory{}
	cf.psqlConnFactory = psql_conn.NewPsqlConnFactory(psqlArgs)
	cf.sgConnFactory = sg_conn.NewSgConnFactory(sgArgs)
	return cf
}

// TODO: Down the line may need to remove this -- not sure as there are more data store relationships
func (f *ControllerFactory) Controller() *BaseController {
	return &BaseController{
		psqlConn: f.psqlConnFactory.Connection(),
		sgConn:   f.sgConnFactory.Connection(),
	}
}

func (f *ControllerFactory) PublicAPIController() PublicAPIController {
	return f.Controller()
}
