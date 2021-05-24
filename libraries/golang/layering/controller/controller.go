package controller

import (
	"web.microservice.shell/libraries/golang/layering/query/psql_conn"
	"web.microservice.shell/libraries/golang/layering/query/sg_conn"
)

// Has all connection types and all methods, interface layer between datastores
type BaseController struct {
	psqlConn *psql_conn.PsqlConnection
	sgConn   *sg_conn.SgConnection
}

func (c *BaseController) PsqlClose() {
	c.psqlConn.Close()
}

func (c *BaseController) SgClose() {
	c.sgConn.Close()
}
