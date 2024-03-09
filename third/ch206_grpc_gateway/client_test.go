package ch206_grpc

import (
	"context"
	pb "go-app/third/ch206_grpc_gateway/hello"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:28080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	in := &pb.HelloRequest{Name: "tom"}
	reply, err := client.SayHello(ctx, in)
	if err != nil {
		panic(err)
	}
	log.Println("send request:", in.Name)
	log.Println("receive response:", reply.GetMessage())
	// Output:
	// 2022/06/12 12:53:54 send request: tom
	// 2022/06/12 12:53:54 receive response: Hello tom
}

// GRPC Gateway HTTP 测试:
// 1.使用两个单独端口
// curl -i 'http://localhost:28081/v1/sayhello' --header 'Content-Type: application/json' --data '{"name": "grpc-gateway"}'
// HTTP/1.1 200 OK
// Content-Type: application/json
// Grpc-Metadata-Content-Type: application/grpc
// Date: Sat, 09 Mar 2024 13:15:38 GMT
// Content-Length: 32
//
// {"message":"Hello grpc-gateway"}%

// 2.使用同一个端口
//  curl -i 'http://localhost:28080/v1/sayhello' --header 'Content-Type: application/json' --data '{"name": "grpc-gateway"}'
// HTTP/1.1 200 OK
// Content-Type: application/json
// Grpc-Metadata-Content-Type: application/grpc
// Grpc-Metadata-Trailer: Grpc-Status
// Grpc-Metadata-Trailer: Grpc-Message
// Grpc-Metadata-Trailer: Grpc-Status-Details-Bin
// Date: Sat, 09 Mar 2024 13:37:14 GMT
// Content-Length: 32
//
// {"message":"Hello grpc-gateway"}
