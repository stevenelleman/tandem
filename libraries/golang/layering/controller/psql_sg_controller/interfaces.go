package psql_sg_controller

import (
	"web.microservice.shell/libraries/golang/layering/models"

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
	CreateSilo(*gin.Context, *models.Silo) error
	UpdateSilo(*gin.Context, *models.Silo) error
	DeleteSilo(*gin.Context, string) error
}

// TODO: Think of better name for controller type that requires SG and PSQL
// Compose methods from two controller
type PublicAPIController interface {
	SgController
	PsqlController
}
