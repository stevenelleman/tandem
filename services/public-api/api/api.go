package main

import (
	"fmt"
	"os"
	"sg/services/public-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// Store host set to sg-api-store
	h := handlers.NewAPIHandler(store, 40)
	defer h.Close()

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
			// Local web-frontend: http://localhost:3000
			// Container web-frontend: http://localhost
			// Tilt: http://localhost:8000
			isLocal := origin == "http://localhost:3000"
			isContainer := origin == "http://localhost"
			isTilt := origin == "http://localhost:8000"
			return isLocal || isContainer || isTilt
		},
	}))

	v1 := router.Group("/v1")
	v1.GET("/silos", h.ListSilos)
	v1.GET("/silos/:silo_id", h.GetSilo)
	v1.POST("/silos/:silo_id", h.CreateSilo)
	v1.PUT("/silos/:silo_id", h.UpdateSilo)
	v1.DELETE("/silos/:silo_id", h.DeleteSilo)

	origin := fmt.Sprintf(":%d", 8000)
	router.Run(origin) // Public API IP:Port
}
