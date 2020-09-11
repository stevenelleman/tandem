package sg_client

import (
	"log"
	"sg/services/public-api/constants"

	pb "libraries/pb/sg"

	grpc "google.golang.org/grpc"
)

type SGClient struct {
	Client *pb.SGClient
}

// TODO: Consider following a controller factory pattern
func (sgc *SGClient) Dial() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(constants.SGServiceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
