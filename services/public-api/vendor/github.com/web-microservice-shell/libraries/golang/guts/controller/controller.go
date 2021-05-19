package controller

import (
	"github.com/web-microservice-shell/libraries/golang/guts/connection/service/psql_conn"
	"github.com/web-microservice-shell/libraries/golang/guts/connection/service/sg_conn"
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
