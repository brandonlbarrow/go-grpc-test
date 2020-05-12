package main

import (
	"fmt"
	"context"
	"go-grpc-test/pkg"
	"google.golang.org/grpc"
	"os"
)

func main() {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		fmt.Printf("Error %s", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := pkg.NewGRPCTestClient(conn)

	message := pkg.MessageInit{
		Body: "I'm the client!",
	}

	response, err := client.GetMessage(context.Background(), &message )
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	fmt.Println(response.GetBody())
}
