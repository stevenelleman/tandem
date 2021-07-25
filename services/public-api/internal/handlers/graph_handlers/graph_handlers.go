package graph_handlers

import (
	"github.com/gin-gonic/gin"
	resp "sg/libraries/golang/layering/responses"
	"sg/services/public-api/internal/controllers/graph_controller"
	reqs "sg/services/public-api/internal/handlers/requests"
)

type GraphHandlers struct {
	controller *graph_controller.GraphController
}

func (h *GraphHandlers) Controller() *graph_controller.GraphController {
	return h.controller
}

func (c GraphHandlers) Close() {
	c.controller.Close()
}

func (h *GraphHandlers) GetNode(ctx *gin.Context) {
	id := ctx.Param("id")
	v, err := h.controller.GetNode(ctx, id)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		resp.ReturnJSON(ctx, 200, v)
	}
}

func (h *GraphHandlers) CreateNode(ctx *gin.Context) {
	req := &reqs.CreateNodeReq{}

	// TODO: make external library for binding to struct
	err := ctx.BindJSON(req)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	}

	err = h.controller.CreateNode(ctx, req.Obj())
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *GraphHandlers) GetEdge(ctx *gin.Context) {
	id := ctx.Param("id")
	v, err := h.controller.GetEdge(ctx, id)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		resp.ReturnJSON(ctx, 200, v)
	}
}

func (h *GraphHandlers) CreateEdge(ctx *gin.Context) {
	req := &reqs.CreateEdgeReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	}

	err = h.controller.CreateEdge(ctx, req.Obj())
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}

func (h *GraphHandlers) GetScope(ctx *gin.Context) {
	id := ctx.Param("id")
	v, err := h.controller.GetScope(ctx, id)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		resp.ReturnJSON(ctx, 200, v)
	}
}

func (h *GraphHandlers) CreateScope(ctx *gin.Context) {
	req := &reqs.CreateScopeReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	}

	err = h.controller.CreateScope(ctx, req.Obj())
	if err != nil {
		resp.ReturnError(ctx, 400, err)
	} else {
		ctx.Status(200)
	}
}


