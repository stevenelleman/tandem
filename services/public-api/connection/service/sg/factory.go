package sg_conn

import (
	"libraries/pb/sg"
	pb "libraries/pb/sg"
	"log"

	"google.golang.org/grpc"
)

type SgConnectionFactory struct {
	conn   *grpc.ClientConn
	client *sg.SGClient
}

func NewSgConnFactory(address string) *SgConnectionFactory {
	sgConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewSGClient(sgConn)

	return &SgConnectionFactory{
		conn:   sgConn,
		client: &client,
	}
}

func (f *SgConnectionFactory) Connection() *SgConnection {
	return &SgConnection{
		conn:   f.conn,
		client: f.client,
	}
}
