package controller

import (
	psql_conn "sg/services/public-api/connection/service/psql"
	sg_conn "sg/services/public-api/connection/service/sg"
	"sg/services/public-api/constants"
)

type ControllerFactory struct {
	psqlConnFactory *psql_conn.PsqlConnectionFactory
	sgConnFactory   *sg_conn.SgConnectionFactory
}

// TODO: Pass in config objects: APIStoreConfig, SGConfig
func NewControllerFactory(host string, conns int) *ControllerFactory {
	cf := &ControllerFactory{}
	cf.psqlConnFactory = psql_conn.NewPsqlConnFactory(host, conns)
	cf.sgConnFactory = sg_conn.NewSgConnFactory(constants.SGServiceAddress)
	return cf
}

func (f *ControllerFactory) Controller() *Controller {
	return &Controller{
		psqlConn: f.psqlConnFactory.Connection(),
		sgConn:   f.sgConnFactory.Connection(),
	}
}
