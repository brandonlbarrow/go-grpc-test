package main

import (
	"context"
	"fmt"
	"go-grpc-test/pkg"
	"google.golang.org/grpc"
	"net"
	"os"
)

type GRPCTestServer struct{}

func (s *GRPCTestServer) GetMessage(c context.Context, m *pkg.MessageInit) (*pkg.MessageReply, error) {
	body := "I'm a server. You said this: " + m.GetBody()
	reply := pkg.MessageReply{
		Body: body,
	}
	return &reply, nil
}

func main() {

	tcp, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error %s", err)
		os.Exit(1)
	}
	srv := grpc.NewServer()
	pkg.RegisterGRPCTestServer(srv, &GRPCTestServer{})
	fmt.Println("Running")
	srv.Serve(tcp)
}
