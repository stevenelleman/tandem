package psql_sg_controller

import (
	"sg/libraries/golang/layering/query/psql_conn"
	"sg/libraries/golang/layering/query/sg_conn"
)

// Has all connection types and all methods, interface layer between datastores
type Controller struct {
	psqlConn *psql_conn.PsqlConnection
	sgConn   *sg_conn.SgConnection
}

func (c *Controller) PsqlClose() {
	c.psqlConn.Close()
}

func (c *Controller) SgClose() {
	c.sgConn.Close()
}
