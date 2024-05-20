package gozeromq_zmq4

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-zeromq/zmq4"
)

func TestPub(t *testing.T) {
	log.SetPrefix("psenvpub: ")

	// prepare the publisher
	pub := zmq4.NewPub(context.Background())
	defer pub.Close()

	err := pub.Listen("tcp://*:5563")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	msgA := zmq4.NewMsgFrom(
		[]byte("A"),
		[]byte("We don't want to see this"),
	)
	msgB := zmq4.NewMsgFrom(
		[]byte("B"),
		[]byte("We would like to see this"),
	)
	for {
		//  Write two messages, each with an envelope and content
		err = pub.Send(msgA)
		if err != nil {
			log.Fatal(err)
		}
		err = pub.Send(msgB)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}

func TestSub(t *testing.T) {
	log.SetPrefix("psenvsub: ")

	//  Prepare our subscriber
	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://localhost:5563")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}

	err = sub.SetOption(zmq4.OptionSubscribe, "B")
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		// Read envelope
		msg, err := sub.Recv()
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Printf("[%s] %s\n", msg.Frames[0], msg.Frames[1])
	}

	// Output:
	// psenvsub: 2022/08/04 12:18:42 [B] We would like to see this
	// psenvsub: 2022/08/04 12:18:43 [B] We would like to see this
	// psenvsub: 2022/08/04 12:18:44 [B] We would like to see this
	// psenvsub: 2022/08/04 12:18:45 [B] We would like to see this
	// psenvsub: 2022/08/04 12:18:46 [B] We would like to see this
	// ...
}
