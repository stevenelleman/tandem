package handlers

import (
	"github.com/gin-gonic/gin"

	"sg/libraries/golang/guts/models"
)

type createSiloRequest struct {
	state string
}

type updateSiloRequest struct {
	state string
}

func (h *BaseHandler) ListSilos(ctx *gin.Context) {
	silos, err := h.PublicAPIController().ListSilos(ctx)
	if err != nil {
		// TODO: Someone is going to forget to put the return statement - how to refactor to avoid?
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSONList(ctx, 200, silos)
	}
}

func (h *BaseHandler) GetSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	silo, err := h.PublicAPIController().GetSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ReturnJSON(ctx, 200, silo)
	}
}

func (h *BaseHandler) CreateSilo(ctx *gin.Context) {
	// TODO: should the id be in the request rather than the path?
	id := ctx.Param("silo_id")
	siloReq := &models.Silo{
		Id: id,
	}

	createReqBody := &createSiloRequest{}
	err := ctx.Bind(createReqBody)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		siloReq.State = createReqBody.state
	}

	err := h.PublicAPIController().CreateSilo(ctx, siloReq)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *BaseHandler) UpdateSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")

	siloReq := &models.Silo{
		Id: id,
	}

	updateReqBody := &updateSiloRequest{}
	err := ctx.Bind(updateReqBody)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		// Must map request object
		siloReq.State = updateReqBody.state
	}

	err := h.PublicAPIController().UpdateSilo(ctx, siloReq)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *BaseHandler) DeleteSilo(ctx *gin.Context) {
	id := ctx.Param("silo_id")
	err := h.PublicAPIController().DeleteSilo(ctx, id)
	if err != nil {
		ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}
