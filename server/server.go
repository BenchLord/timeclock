package main

import (
	"context"

	pb "github.com/benchlord/timeclock/protos"
)

// TimeClockServer ...
type TimeClockServer struct{}

// ClockIn ...
func (t *TimeClockServer) ClockIn(ctx context.Context, req *pb.ClockInRequest) (*pb.ClockInResponse, error) {
	return &pb.ClockInResponse{
		Id: 1,
	}, nil
}

// ClockOut ...
func (t *TimeClockServer) ClockOut(ctx context.Context, req *pb.ClockOutRequest) (*pb.ClockOutResponse, error) {
	return &pb.ClockOutResponse{}, nil
}

// GetHours ...
func (t *TimeClockServer) GetHours(ctx context.Context, req *pb.GetHoursRequest) (*pb.GetHoursResponse, error) {
	return &pb.GetHoursResponse{}, nil
}

