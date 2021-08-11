package sg_conn

import (
	"web.microservice.shell/libraries/golang/pb/sg"

	"google.golang.org/grpc"
)

type SgConnection struct {
	conn   *grpc.ClientConn
	client *sg.SGClient
}

func (c *SgConnection) Close() {
	c.conn.Close()
}
