package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	"github.com/benchlord/timeclock/protos/auth"
)

// func main() {
// 	cmd.Execute()
// }

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	client := auth.NewAuthServiceClient(conn)
	ctx := context.Background()
	res, err := client.Register(ctx, &auth.RegisterRequest{
		Username: "brandon",
		Password: "pass",
	})
	fmt.Printf("Res: %+v\nErr: %+v\n", res, err)
}
