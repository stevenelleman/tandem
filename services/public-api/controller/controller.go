package controller

import (
	"sg/services/public-api/connection"
	sg_client "sg/services/public-api/controller/sgclient"
)

type Controller struct {
	conn     *connection.Connection
	sgClient *sg_client.SGClient
}

func (c Controller) Close() {
	c.conn.Close()
}
