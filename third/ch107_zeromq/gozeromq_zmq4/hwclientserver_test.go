package gozeromq_zmq4

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	zmq "github.com/go-zeromq/zmq4"
)

func TestHWServer(t *testing.T) {
	if err := hwserver(); err != nil {
		log.Fatalf("hwserver: %v", err)
	}
	// Output:
	// Received  Msg{Frames:{"hello-0"}} hello-0
	// Received  Msg{Frames:{"hello-1"}} hello-1
	// Received  Msg{Frames:{"hello-2"}} hello-2
}

func hwserver() error {
	ctx := context.Background()
	// Socket to talk to clients
	socket := zmq.NewRep(ctx)
	defer socket.Close()
	if err := socket.Listen("tcp://*:5555"); err != nil {
		return fmt.Errorf("listening: %v", err)
	}

	for {
		msg, err := socket.Recv()
		if err != nil {
			return fmt.Errorf("receiving: %w", err)
		}
		fmt.Println("Received ", msg, string(msg.Bytes()))

		// Do some 'work'
		time.Sleep(time.Second)

		reply := fmt.Sprintf("World" + strings.Trim(string(msg.Bytes()), "hello"))
		if err := socket.Send(zmq.NewMsgString(reply)); err != nil {
			return fmt.Errorf("sending reply: %w", err)
		}
	}
}

func TestHWClient(t *testing.T) {
	if err := hwclient(); err != nil {
		log.Fatalf("hwclient: %v", err)
	}
	// Output:
	// Connecting to hello world server...
	// sending  Msg{Frames:{"hello-0"}}
	// received  Msg{Frames:{"World-0"}}
	// sending  Msg{Frames:{"hello-1"}}
	// received  Msg{Frames:{"World-1"}}
	// sending  Msg{Frames:{"hello-2"}}
	// received  Msg{Frames:{"World-2"}}
}

func hwclient() error {
	ctx := context.Background()
	socket := zmq.NewReq(ctx, zmq.WithDialerRetry(time.Second))
	defer socket.Close()

	fmt.Println("Connecting to hello world server...")
	if err := socket.Dial("tcp://localhost:5555"); err != nil {
		return fmt.Errorf("dialing: %w", err)
	}

	for i := 0; i < 3; i++ {
		// Send hello.
		m := zmq.NewMsgString(fmt.Sprintf("%s-%d", "hello", i))
		fmt.Println("sending ", m)
		if err := socket.Send(m); err != nil {
			return fmt.Errorf("sending: %w", err)
		}

		// Wait for reply.
		r, err := socket.Recv()
		if err != nil {
			return fmt.Errorf("receiving: %w", err)
		}
		fmt.Println("received ", r.String())
	}
	return nil
}
