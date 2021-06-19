package main

import (
	"github.com/gin-gonic/gin"
	"web.microservice.shell/libraries/golang/layering/handler/psql_sg_handler"
	"web.microservice.shell/services/public-api/args"
	"os"
)

func main() {
	host := os.Args[1]
	if host == "" {
		panic("Host not defined")
	}
	store := os.Args[2]
	if store == "" {
		panic("Store not defined")
	}

	psqlArgs, sgArgs := args.MakeArgsFromEnv(store)
	h := psql_sg_handler.NewPublicAPIHandler(psqlArgs, sgArgs)
	defer h.PublicAPIClose()

	// TODO: need authz middleware to convert cookie into faceted identity list.
	// 	The faceted identity will be associated with a list of silos
	// 	One or more FIs must belong to the subject silo
	//	One or more Silos must belong to the subject forum
	//	How to avoid leakage? Only want to grab relevant FI, not everything. Pass in FI list and subject forum/silo
	//	And confirm an overlap

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/silos", h.ListSilos)
	v1.GET("/silos/:silo_id", h.GetSilo)
	v1.POST("/silos/:silo_id", h.CreateSilo)
	v1.PUT("/silos/:silo_id", h.UpdateSilo)
	v1.DELETE("/silos/:silo_id", h.DeleteSilo)

	router.Run(os.Getenv("PUBLIC_API_ADDRESS")) // Public API IP:Port
}
