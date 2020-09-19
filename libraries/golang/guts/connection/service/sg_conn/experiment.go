package sg_conn

import (
	"context"
	"fmt"
	pb "sg/libraries/golang/pb/sg"
	"time"

	"github.com/gin-gonic/gin"
)

// Dummy code to demonstrate that the grpc connection works
func (c *SgConnection) SayHello(ctx *gin.Context) error {
	reqctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for i := range [10]int{} {
		// TODO: Is this a terrible pattern?
		reply, err := (*c.client).SayHello(reqctx, &pb.HelloRequest{Name: "BLAHBLAHBLAH"})
		if err != nil {
			return err
		}
		fmt.Println("Reply", i, reply.GetMessage())
	}

	return nil
}
