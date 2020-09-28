package anywhat

import (
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	Service pb.AnywhatClient
}

// NewClient ...
func NewClient(url string) *Client {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := pb.NewAnywhatClient(conn)
	return &Client{conn, c}
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}
