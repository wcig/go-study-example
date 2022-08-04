package gozeromq_zmq4

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-zeromq/zmq4"
)

func TestRRClient(t *testing.T) {
	log.SetPrefix("rrclient: ")

	req := zmq4.NewReq(context.Background())
	defer req.Close()

	err := req.Dial("tcp://localhost:5559")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}

	for i := 0; i < 3; i++ {
		err := req.Send(zmq4.NewMsgString("Hello"))
		if err != nil {
			log.Fatalf("could not send greeting: %v", err)
		}

		msg, err := req.Recv()
		if err != nil {
			log.Fatalf("could not recv greeting: %v", err)
		}
		log.Printf("received reply %d [%s]\n", i, msg.Frames[0])
	}

	// Output:
	// rrclient: 2022/08/04 12:25:28 received reply 0 [World]
	// rrclient: 2022/08/04 12:25:29 received reply 1 [World]
	// rrclient: 2022/08/04 12:25:30 received reply 2 [World]
}

func TestRRWorker(t *testing.T) {
	log.SetPrefix("rrworker: ")

	//  Socket to talk to clients
	rep := zmq4.NewRep(context.Background())
	defer rep.Close()

	err := rep.Listen("tcp://*:5559")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}

	for {
		//  Wait for next request from client
		msg, err := rep.Recv()
		if err != nil {
			log.Fatalf("could not recv request: %v", err)
		}

		log.Printf("received request: [%s]\n", msg.Frames[0])

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		err = rep.Send(zmq4.NewMsgString("World"))
		if err != nil {
			log.Fatalf("could not send reply: %v", err)
		}
	}

	// Output:
	// rrworker: 2022/08/04 12:25:27 received request: [Hello]
	// rrworker: 2022/08/04 12:25:28 received request: [Hello]
	// rrworker: 2022/08/04 12:25:29 received request: [Hello]
}
