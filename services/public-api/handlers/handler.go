package handlers

import "sg/services/public-api/controller"

type APIHandler struct {
	Controller *controller.Controller
}

func (c APIHandler) Close() {
	c.Controller.Close()
}
