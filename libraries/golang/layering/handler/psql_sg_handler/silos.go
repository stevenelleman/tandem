package psql_sg_handler

import (
	"github.com/gin-gonic/gin"

	"web.microservice.shell/libraries/golang/layering/models"
)

func (h *Handler) ListSilos(ctx *gin.Context) {
	silos, err := h.Controller().ListSilos(ctx)
	if err != nil {
		// TODO: Someone is going to forget to put the return statement - how to refactor to avoid?
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSONList(ctx, 200, silos)
	}
}

func (h *Handler) GetSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	silo, err := h.Controller().GetSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSON(ctx, 200, silo)
	}
}

func (h *Handler) CreateSilo(ctx *gin.Context) {
	// TODO: should the id be in the request rather than the path?
	id := ctx.Param("silo_id")
	silo := &models.Silo{
		Id: id,
	}

	req := &createSiloReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		silo.State = req.State
	}

	err = h.Controller().CreateSilo(ctx, silo)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *Handler) UpdateSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	silo := &models.Silo{
		Id: id,
	}

	req := &updateSiloReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		silo.State = req.State
	}

	err = h.Controller().UpdateSilo(ctx, silo)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *Handler) DeleteSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	err := h.Controller().DeleteSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}
