package main

import (
	"context"
	"fmt"
	pb "github.com/web-microservice-shell/libraries/golang/pb/sg"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/web-microservice-shell/services/grpc/constants"
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
