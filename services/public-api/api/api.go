package main

import (
	"github.com/gin-gonic/gin"

	"sg/services/public-api/db"
	"sg/services/public-api/handlers"
)

func main() {
	db.InitDb()
	defer db.DB.Close()

	// TODO: need authz middleware to convert cookie into FI list
	router := gin.Default()

	// TODO: Also pass in some kind of response writer object
	// Let the API endpoints begin...
	v1 := router.Group("/v1")
	v1.GET("/silos", handlers.ListSilos) // Works
	v1.GET("/silos/:silo_id", handlers.GetSilo) // Works
	v1.POST("/silos/:silo_id", handlers.CreateSilo) // Works
	v1.PUT("/silos/:silo_id", handlers.UpdateSilo)
	v1.DELETE("/silos/:silo_id", handlers.DeleteSilo) // Works

	router.Run("localhost:8000")
}
