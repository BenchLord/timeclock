package auth

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/benchlord/timeclock/protos/auth"
)

type AuthHandler struct {
	Db *sql.DB
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// hash password
	res, err := h.Db.ExecContext(ctx, `insert into users (username, password) values ($1, $2)`, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &pb.RegisterResponse{}, nil
}
