package sg_client

import (
	"log"

	pb "libraries/pb/sg"

	grpc "google.golang.org/grpc"
)

const (
	address = "172.18.0.5:50051"
)

type SGClient struct {
	Client *pb.SGClient
}

func (sgc *SGClient) Dial() *grpc.ClientConn {
	// Set up a connection to the server.
	// grpc.WithInsecure(),
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
	//defer conn.Close()

	// c := pb.NewSGClient(conn)
	// sgc.Client = &c

	/*ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "dfgsfgsgsdfgdsg"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())*/
}
