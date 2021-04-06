package sg_conn

import (
	"google.golang.org/grpc"
	sg2 "libraries/golang/pb/sg"
	"log"
)

type StoreArgs struct {
	address string
}

func MakeArgs(address string) *StoreArgs {
	return &StoreArgs{
		address: address,
	}
}

func (a *StoreArgs) Address() string {
	return a.address
}

type SgConnectionFactory struct {
	conn   *grpc.ClientConn
	client *sg2.SGClient
}

func NewSgConnFactory(args *StoreArgs) *SgConnectionFactory {
	sgConn, err := grpc.Dial(args.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := sg2.NewSGClient(sgConn)

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
