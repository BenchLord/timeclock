package main

import (
	"net"

	pb "github.com/benchlord/timeclock/protos"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	pb.RegisterTimeClockServer(server, &TimeClockServer{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
