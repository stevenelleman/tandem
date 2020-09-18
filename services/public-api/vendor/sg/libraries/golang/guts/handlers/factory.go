package handlers

import (
	"sg/libraries/golang/guts/connection/service/psql_conn"
	"sg/libraries/golang/guts/connection/service/sg_conn"

	"sg/libraries/golang/guts/controller"

	"github.com/gin-gonic/gin"
)

type PublicAPIHandler interface {
	Close()
	Controller() *controller.Controller // Controller specific to PublicAPI

	ListSilos(ctx *gin.Context)
	GetSilo(ctx *gin.Context)
	CreateSilo(ctx *gin.Context)
	UpdateSilo(ctx *gin.Context)
	DeleteSilo(ctx *gin.Context)
}

func NewPublicAPIHandler(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) PublicAPIHandler {
	cf := controller.NewControllerFactory(psqlArgs, sgArgs)
	return &Handler{
		controller: cf.Controller(),
	}
}
