package ch206_grpc

import (
	"context"
	pb "go-app/third/ch206_grpc_gateway/hello"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcServerEndpoint = "localhost:28080"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func TestServer(t *testing.T) {
	go TestGPRCServer(t)
	TestGRPCGatewayHTTPServer(t)
}

func TestGPRCServer(t *testing.T) {
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

func TestGRPCGatewayHTTPServer(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":28081", mux); err != nil {
		log.Fatal(err)
	}
}
