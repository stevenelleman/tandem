package handlers

import (
	"github.com/gin-gonic/gin"

	"sg/services/public-api/handlers/requests"
)

func (h *APIHandler) ListSilos(ctx *gin.Context) {
	silos, err := h.Controller.ListSilos(ctx)
	if err != nil {
		// TODO: Someone is going to forget to put the return statement - how to refactor to avoid?
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSONList(ctx, 200, silos)
	}
}

func (h *APIHandler) GetSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	silo, err := h.Controller.GetSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSON(ctx, 200, silo)
	}
}

func (h *APIHandler) CreateSilo(ctx *gin.Context) {
	// TODO: should the id be in the request rather than the path?
	id := ctx.Param("silo_id")

	// Must map request object
	siloReq := &requests.Silo{
		Id: id,
	}

	err := h.Controller.CreateSilo(ctx, siloReq)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *APIHandler) UpdateSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")

	// Must map request object
	siloReq := &requests.Silo{
		Id: id,
	}

	err := h.Controller.UpdateSilo(ctx, siloReq)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *APIHandler) DeleteSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	err := h.Controller.DeleteSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}
