package handlers

import "github.com/web-microservice-shell/libraries/golang/guts/controller"

type BaseHandler struct {
	controller *controller.BaseController
}

func (h *BaseHandler) PublicAPIController() controller.PublicAPIController {
	return h.controller
}

func (h *BaseHandler) Controller() *controller.BaseController {
	return h.controller
}

func (c BaseHandler) PublicAPIClose() {
	c.controller.PsqlClose()
	c.controller.SgClose()
}
