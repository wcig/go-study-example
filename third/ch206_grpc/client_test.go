package ch206_grpc

import (
	"context"
	pb "go-app/third/ch206_grpc/helloworld"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
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
