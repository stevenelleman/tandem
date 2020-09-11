package controller

import (
	"context"
	"fmt"
	pb "libraries/pb/sg"
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ListSilos(ctx *gin.Context) ([]*models.Silo, error) {
	// TODO: Clean up pattern
	reqctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Dummy code to demonstrate that the grpc connection works
	// The addition of this client begins to specify/"fix" this pattern
	conn := c.sgClient.Dial()
	defer conn.Close()

	client := pb.NewSGClient(conn)
	for i := range [10]int{} {
		reply, err := client.SayHello(reqctx, &pb.HelloRequest{Name: "BLAHBLAHBLAH"})
		if err != nil {
			return nil, err
		}
		fmt.Println("Reply", i, reply.GetMessage())
	}

	return c.conn.ListSilos(ctx)
}

func (c *Controller) GetSilo(ctx *gin.Context, id string) (*models.Silo, error) {
	return c.conn.GetSilo(ctx, id)
}

func (c *Controller) CreateSilo(ctx *gin.Context, s *requests.Silo) error {
	return c.conn.CreateSilo(ctx, s)
}

func (c *Controller) UpdateSilo(ctx *gin.Context, s *requests.Silo) error {
	return c.conn.UpdateSilo(ctx, s)
}

func (c *Controller) DeleteSilo(ctx *gin.Context, id string) error {
	return c.conn.DeleteSilo(ctx, id)
}
