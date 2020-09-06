package controller

import "sg/services/public-api/connection"

type Controller struct {
	Connection *connection.Connection
}

func (c Controller) Close() {
	c.Connection.Close()
}
