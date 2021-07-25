package psql_sg_handler

import (
	"sg/libraries/golang/layering/controller/psql_sg_controller"
)

type Handler struct {
	controller *psql_sg_controller.Controller
}

func (h *Handler) Controller() *psql_sg_controller.Controller {
	return h.controller
}

func (c Handler) Close() {
	c.controller.PsqlClose()
	c.controller.SgClose()
}
