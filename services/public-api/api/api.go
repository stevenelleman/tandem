package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shell/libraries/golang/guts/handlers"
	"github.com/shell/services/public-api/args"
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

	// Store host set to api-store
	// TODO: Re-use in other golang services -- all it should take is passing in store info and it should work
	psqlArgs, sgArgs := args.MakeArgsFromEnv(store)

	h := handlers.NewPublicAPIHandler(psqlArgs, sgArgs)
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
