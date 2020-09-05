package main

import (
	"fmt"
	"os"
	"sg/services/public-api/db"
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

	db.InitDb(host)
	defer db.DB.Close()

	// TODO: need authz middleware to convert cookie into FI list
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			fmt.Println("Allowed origin", origin)

			// Local web-frontend: http://localhost:3000
			// Container web-frontend: http://localhost
			isLocal := origin == "http://localhost:3000"
			isContainer := origin == "http://localhost"

			return isLocal || isContainer
		},
	}))

	// TODO: Also pass in some kind of response writer object
	// Let the API endpoints begin...
	v1 := router.Group("/v1")
	v1.GET("/silos", handlers.ListSilos)
	v1.GET("/silos/:silo_id", handlers.GetSilo)
	v1.POST("/silos/:silo_id", handlers.CreateSilo)
	v1.PUT("/silos/:silo_id", handlers.UpdateSilo)
	v1.DELETE("/silos/:silo_id", handlers.DeleteSilo)

	origin := fmt.Sprintf("%s:%d", ip, 8000)
	router.Run(origin) // Public API IP:Port
}
