package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"

	"github.com/benchlord/timeclock/server/auth"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	db, err := sql.Open("postgres", "sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	server := grpc.NewServer()
	authService := &auth.AuthService{
		Handler: &auth.AuthHandler{
			Db: db,
		},
	}
	authService.Register(server)

	if err := server.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
