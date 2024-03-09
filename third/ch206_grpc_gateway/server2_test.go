package ch206_grpc

import (
	"context"
	pb "go-app/third/ch206_grpc_gateway/hello"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

// GRPC server 和 GRPC Gateway HTTP server 使用同一个断端口
func TestServerWithOnePort(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:28080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Server{})

	// log.Printf("server listening at %v", listener.Addr())
	// if err = server.Serve(listener); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatal(err)
	}

	otherHandler := http.NewServeMux()
	otherHandler.Handle("/", mux)
	gwServer := &http.Server{
		Addr:    "localhost:28080",
		Handler: grpcHandlerFunc(server, otherHandler),
	}
	if err = gwServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
	server := &http2.Server{}
	return h2c.NewHandler(handler, server)
}
