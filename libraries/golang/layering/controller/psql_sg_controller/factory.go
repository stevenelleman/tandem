package psql_sg_controller

import (
	"web.microservice.shell/libraries/golang/layering/query/psql_conn"
	"web.microservice.shell/libraries/golang/layering/query/sg_conn"
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

func (f *ControllerFactory) Controller() *Controller {
	return &Controller{
		psqlConn: f.psqlConnFactory.Connection(),
		sgConn:   f.sgConnFactory.Connection(),
	}
}
