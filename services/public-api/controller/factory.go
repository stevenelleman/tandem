package controller

import (
	"sg/services/public-api/connection"
	sg_client "sg/services/public-api/controller/sgclient"
)

type ControllerFactory struct {
	connectionFactory *connection.ConnectionFactory
	sgClient          *sg_client.SGClient
}

func NewControllerFactory(host string, conns int) *ControllerFactory {
	cf := &ControllerFactory{}
	// Can also consider passing the connection factory
	cf.connectionFactory = connection.NewConnectionFactory(host, conns)

	cf.sgClient = &sg_client.SGClient{}
	return cf
}

func (f *ControllerFactory) Controller() *Controller {
	return &Controller{
		conn:     f.connectionFactory.Connection(),
		sgClient: f.sgClient,
	}
}
