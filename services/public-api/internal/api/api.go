package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"sg/services/public-api/internal/args"
	"sg/services/public-api/internal/handlers/graph_handlers"
)

func Run(store string) {
	psqlArgs := args.MakeArgsFromEnv(store)
	h := graph_handlers.NewHandler(psqlArgs)
	defer h.Close()

	router := gin.Default()
	v1 := router.Group("/v1")

	// NOTE: May be more appropriate to use web sockets rather than REST?
	v1.POST("/node", h.CreateNode)
	v1.GET("/node/:id", h.GetNode)
	v1.POST("/edge", h.CreateEdge)
	v1.GET("/edge/:id", h.GetEdge)
	v1.POST("/scope", h.CreateScope)
	v1.GET("/scope/:id", h.GetScope)

	router.Run(os.Getenv("PUBLIC_API_ADDRESS")) // Public API IP:Port
}
