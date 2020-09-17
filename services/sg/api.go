package main

import (
	"context"
	"fmt"
	pb "libraries/pb/sg"
	"log"
	"net"
	"sg/services/sourcegraph/constants"

	"google.golang.org/grpc"
)

func sayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	fmt.Printf("Listen on %s\n", constants.Port)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// TODO: How to image as part of handler-controller-connection pattern?
	s := grpc.NewServer()
	pb.RegisterSGService(s, &pb.SGService{SayHello: sayHello})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
