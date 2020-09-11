package controller

import (
	"context"
	"fmt"
	pb "libraries/pb/sg"
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
	"time"
)

func (c *Controller) ListSilos() ([]*models.Silo, error) {
	// TODO: Clean up pattern
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Dummy code to demonstrate that the grpc connection works
	conn := c.sgClient.Dial()
	defer conn.Close()

	client := pb.NewSGClient(conn)
	for i := range [10]int{} {
		reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "BLAHBLAHBLAH"})
		if err != nil {
			return nil, err
		}
		fmt.Println("Reply", i, reply.GetMessage())
	}

	return c.conn.ListSilos()
}

func (c *Controller) GetSilo(id string) (*models.Silo, error) {
	return c.conn.GetSilo(id)
}

func (c *Controller) CreateSilo(s *requests.Silo) error {
	return c.conn.CreateSilo(s)
}

func (c *Controller) UpdateSilo(s *requests.Silo) error {
	return c.conn.UpdateSilo(s)
}

func (c *Controller) DeleteSilo(id string) error {
	return c.conn.DeleteSilo(id)
}
