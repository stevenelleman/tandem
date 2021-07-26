package main

import (
	"context"
	pb "sg/libraries/golang/pb/sg"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func sayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	address := os.Getenv("GRPC_ADDRESS")
	if address == "" {
		panic("Port environment variable not defined")
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// TODO: How to imagine as part of handler-controller-connection pattern?
	s := grpc.NewServer()
	pb.RegisterSGService(s, &pb.SGService{SayHello: sayHello})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
