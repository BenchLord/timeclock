package dialer

import (
	pb "github.com/benchlord/timeclock/protos"
	"google.golang.org/grpc"
)

// TODO: change package
// probably add to a new dir 'client'

// Client ...
type Client struct {
	grpc.ClientConn
	pb.TimeClockClient
}

// NewClient ...
func NewClient() (*Client, error) {
	// TODO: get grpc client
	return &Client{}, nil
}
