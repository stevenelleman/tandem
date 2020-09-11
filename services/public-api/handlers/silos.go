package handlers

import (
	"github.com/gin-gonic/gin"

	"sg/services/public-api/handlers/requests"
)

func (h *APIHandler) ListSilos(c *gin.Context) {
	silos, err := h.Controller.ListSilos()
	if err != nil {
		// TODO: Someone is going to forget to put the return statement - how to refactor to avoid?
		ReturnError(c, 400, err)
	} else {
		ReturnJSONList(c, 200, silos)
	}
}

func (h *APIHandler) GetSilo(c *gin.Context) {
	id := c.Param("silo_id")
	silo, err := h.Controller.GetSilo(id)
	if err != nil {
		ReturnError(c, 400, err)
	} else {
		ReturnJSON(c, 200, silo)
	}
}

func (h *APIHandler) CreateSilo(c *gin.Context) {
	// TODO: should the id be in the request rather than the path?
	id := c.Param("silo_id")

	// Must map request object
	siloReq := &requests.Silo{
		Id: id,
	}

	err := h.Controller.CreateSilo(siloReq)
	if err != nil {
		ReturnError(c, 400, err)
	} else {
		c.Status(200)
	}
}

func (h *APIHandler) UpdateSilo(c *gin.Context) {
	id := c.Param("silo_id")

	// Must map request object
	siloReq := &requests.Silo{
		Id: id,
	}

	err := h.Controller.UpdateSilo(siloReq)
	if err != nil {
		ReturnError(c, 400, err)
	} else {
		c.Status(200)
	}
}

func (h *APIHandler) DeleteSilo(c *gin.Context) {
	id := c.Param("silo_id")
	err := h.Controller.DeleteSilo(id)
	if err != nil {
		ReturnError(c, 400, err)
	} else {
		c.Status(200)
	}
}
