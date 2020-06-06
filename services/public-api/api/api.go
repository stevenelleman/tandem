package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"sg/services/public-api/db"
	"sg/services/public-api/handlers"
)

func main() {
	db.InitDb()
	defer db.DB.Close()

	// TODO: need authz middleware to convert cookie into FI list
	router := gin.Default()

	// TODO: Pass in as argument
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
	}))

	// TODO: Also pass in some kind of response writer object
	// Let the API endpoints begin...
	v1 := router.Group("/v1")
	v1.GET("/silos", handlers.ListSilos)
	v1.GET("/silos/:silo_id", handlers.GetSilo)
	v1.POST("/silos/:silo_id", handlers.CreateSilo)
	v1.PUT("/silos/:silo_id", handlers.UpdateSilo)
	v1.DELETE("/silos/:silo_id", handlers.DeleteSilo)

	// TODO: Pass in as argument
	router.Run("localhost:8000")
}
