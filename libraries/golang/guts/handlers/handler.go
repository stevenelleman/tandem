package handlers

import "sg/libraries/golang/guts/controller"

type Handler struct {
	controller *controller.Controller
}

func (c Handler) Close() {
	c.controller.Close()
}
