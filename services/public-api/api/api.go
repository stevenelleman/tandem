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
	ip := os.Args[1]
	if ip == "" {
		panic("IP not defined")
	}

	host := "172.18.0.3" // Postgres IP:Port
	if ip == "localhost" {
		// If host is localhost then use `stevenelleman` db
		host = ip
	}

	h := handlers.NewAPIHandler(host, 40)
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
			// Local web-frontend: http://localhost:3000
			// Container web-frontend: http://localhost
			isLocal := origin == "http://localhost:3000"
			isContainer := origin == "http://localhost"

			return isLocal || isContainer
		},
	}))

	v1 := router.Group("/v1")
	v1.GET("/silos", h.ListSilos)
	v1.GET("/silos/:silo_id", h.GetSilo)
	v1.POST("/silos/:silo_id", h.CreateSilo)
	v1.PUT("/silos/:silo_id", h.UpdateSilo)
	v1.DELETE("/silos/:silo_id", h.DeleteSilo)

	origin := fmt.Sprintf("%s:%d", ip, 8000)
	router.Run(origin) // Public API IP:Port
}
