package sg_conn

import (
	"github.com/web-microservice-shell/libraries/golang/pb/sg"

	"google.golang.org/grpc"
)

type SgConnection struct {
	conn   *grpc.ClientConn
	client *sg.SGClient
}

func (c *SgConnection) Close() {
	c.conn.Close()
}
