package psql_sg_handler

import (
	"web.microservice.shell/libraries/golang/layering/controller/psql_sg_controller"
)

type BaseHandler struct {
	controller *psql_sg_controller.BaseController
}

func (h *BaseHandler) PublicAPIController() psql_sg_controller.PublicAPIController {
	return h.controller
}

func (h *BaseHandler) Controller() *psql_sg_controller.BaseController {
	return h.controller
}

func (c BaseHandler) PublicAPIClose() {
	c.controller.PsqlClose()
	c.controller.SgClose()
}
