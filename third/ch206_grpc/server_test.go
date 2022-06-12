package ch206_grpc

import (
	"context"
	pb "go-app/third/ch206_grpc/helloworld"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func TestServer(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:28080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Server{})
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// Output:
	// 2022/06/12 12:52:13 server listening at 127.0.0.1:28080
	// 2022/06/12 12:53:54 Received: tom
}
