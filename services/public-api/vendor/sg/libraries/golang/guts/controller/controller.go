package controller

import (
	"sg/libraries/golang/guts/connection/service/psql_conn"
	"sg/libraries/golang/guts/connection/service/sg_conn"
)

type Controller struct {
	psqlConn *psql_conn.PsqlConnection
	sgConn   *sg_conn.SgConnection
}

// HERE: would likely be a problem if the connection was never opened
func (c Controller) Close() {
	c.PsqlClose()
	c.SgClose()
}
