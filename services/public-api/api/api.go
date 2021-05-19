package main

import (
	"fmt"
	"os"
	"github.com/shell/libraries/golang/datastore/migrater"
	"github.com/shell/libraries/golang/guts/connection/service/psql_conn"
	"github.com/shell/libraries/golang/guts/connection/service/sg_conn"
	"github.com/shell/libraries/golang/guts/handlers"
	"github.com/shell/services/public-api/constants"
	"github.com/shell/services/public-api/mapper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: Use same structure in sg service, but a grpc router instead of a rest one
// 	Followed by handler which returns controller object into a specific interface, controller object will be dependent on data stores
// 	and other connected services
func main() {
	// Pass in target IP of Public API service. Localhost for individually run, 172.17.0.2 for container
	host := os.Args[1]
	if host == "" {
		panic("Host not defined")
	}
	store := os.Args[2]
	if store == "" {
		panic("Store not defined")
	}
	// Store host set to api-store
	// TODO: Re-use in other golang services -- all it should take is passing in store info and it should work

	migraterArgs := migrater.MakeArgs(
		constants.APIStoreDriver,
		constants.APIStoreMigraterTableName,
		constants.APIStoreMigrationPath,
	)

	psqlArgs := psql_conn.MakeArgs(
		constants.APIStoreDriver,
		store,
		constants.APIStoreUser,
		constants.APIStorePassword,
		constants.APIStoreName,
		constants.APIStorePort,
		constants.MaxConns,
		migraterArgs,
		mapper.MapTables,
	)

	sgArgs := sg_conn.MakeArgs(constants.SGServiceAddress)

	h := handlers.NewPublicAPIHandler(psqlArgs, sgArgs)
	defer h.PublicAPIClose()

	// TODO: need authz middleware to convert cookie into faceted identity list.
	// 	The faceted identity will be associated with a list of silos
	// 	One or more FIs must belong to the subject silo
	//	One or more Silos must belong to the subject forum
	//	How to avoid leakage? Only want to grab relevant FI, not everything. Pass in FI list and subject forum/silo
	//	And confirm an overlap

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// TODO: Must change origin
			isLocal := origin == constants.LocalWebFrontendHost
			isContainer := origin == constants.DockerComposeHost
			isTilt := origin == constants.TiltHost
			return isLocal || isContainer || isTilt
		},
	}))

	v1 := router.Group("/v1")
	v1.GET("/silos", h.ListSilos)
	v1.GET("/silos/:silo_id", h.GetSilo)
	v1.POST("/silos/:silo_id", h.CreateSilo)
	v1.PUT("/silos/:silo_id", h.UpdateSilo)
	v1.DELETE("/silos/:silo_id", h.DeleteSilo)

	origin := fmt.Sprintf(":%d", constants.PublicAPIPort)
	router.Run(origin) // Public API IP:Port
}
