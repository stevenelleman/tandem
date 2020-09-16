package controller

import (
	psql_conn "sg/services/public-api/connection/service/psql"
	sg_conn "sg/services/public-api/connection/service/sg"
)

type Controller struct {
	psqlConn *psql_conn.PsqlConnection
	sgConn   *sg_conn.SgConnection
}

func (c Controller) Close() {
	c.psqlConn.Close()
	c.sgConn.Close()
}
