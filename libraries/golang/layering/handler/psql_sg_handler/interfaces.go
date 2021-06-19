package psql_sg_handler

import (
	"web.microservice.shell/libraries/golang/layering/controller/psql_sg_controller"

	"github.com/gin-gonic/gin"
)

// Handler interface for Public-API Service
type PublicAPIHandler interface {
	// Closes PsqlController and SgController
	PublicAPIClose()

	// Should be the Controller specific to PublicAPI and used in methods
	PublicAPIController() psql_sg_controller.PublicAPIController

	ListSilos(ctx *gin.Context)
	GetSilo(ctx *gin.Context)
	CreateSilo(ctx *gin.Context)
	UpdateSilo(ctx *gin.Context)
	DeleteSilo(ctx *gin.Context)
}
