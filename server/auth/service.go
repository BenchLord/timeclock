package auth

import (
	"google.golang.org/grpc"

	pb "github.com/benchlord/timeclock/protos/auth"
)

type AuthService struct {
	Handler *AuthHandler
}

func (a *AuthService) Register(server *grpc.Server) error {
	pb.RegisterAuthServiceServer(server, a.Handler)
	return nil
}

func (a *AuthService) Close() error {
	if a.Handler.Db != nil {
		if err := a.Handler.Db.Close(); err != nil {
			return err
		}
		a.Handler.Db = nil
	}
	return nil
}
