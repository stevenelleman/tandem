package handlers

import (
	"web-microservice-shell/libraries/golang/guts/controller"

	"github.com/gin-gonic/gin"
)

// Handler interface for Public-API Service
type PublicAPIHandler interface {
	// Closes PsqlController and SgController
	PublicAPIClose()

	// Should be the Controller specific to PublicAPI and used in methods
	PublicAPIController() controller.PublicAPIController

	ListSilos(ctx *gin.Context)
	GetSilo(ctx *gin.Context)
	CreateSilo(ctx *gin.Context)
	UpdateSilo(ctx *gin.Context)
	DeleteSilo(ctx *gin.Context)
}
