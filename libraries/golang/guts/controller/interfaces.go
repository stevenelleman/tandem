package controller

import (
	"sg/libraries/golang/guts/handlers/requests"
	"sg/libraries/golang/guts/models"

	"github.com/gin-gonic/gin"
)

type SgController interface {
	SgClose()

	SayHello(*gin.Context) error
}

type PsqlController interface {
	PsqlClose()

	ListSilos(*gin.Context) ([]*models.Silo, error)
	GetSilo(*gin.Context, string) (*models.Silo, error)
	CreateSilo(*gin.Context, *requests.Silo) error
	UpdateSilo(*gin.Context, *requests.Silo) error
	DeleteSilo(*gin.Context, string) error
}

// TODO: Think of better name for controller type that requires SG and PSQL
// Compose methods from two controller
type PublicAPIController interface {
	SgController
	PsqlController
}
