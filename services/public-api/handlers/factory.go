package handlers

import (
	"sg/services/public-api/controller"
)

func NewAPIHandler(host string, conns int) *APIHandler {
	cf := controller.NewControllerFactory(host, conns)
	return &APIHandler{
		Controller: cf.Controller(),
	}
}
