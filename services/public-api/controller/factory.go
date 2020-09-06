package controller

import (
	"sg/services/public-api/connection"
)

type ControllerFactory struct {
	connectionFactory *connection.ConnectionFactory
}

func NewControllerFactory(host string, conns int) *ControllerFactory {
	cf := &ControllerFactory{}
	// Can also consider passing the connection factory
	cf.connectionFactory = connection.NewConnectionFactory(host, conns)
	return cf
}

func (f *ControllerFactory) Controller() *Controller {
	return &Controller{
		Connection: f.connectionFactory.Connection(),
	}
}
