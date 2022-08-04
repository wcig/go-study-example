package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:5555")

	// Do 10 requests, waiting each time for a response
	for i := 0; i < 10; i++ {
		fmt.Printf("Sending request %d...\n", i)
		s.Send("Hello", 0)

		msg, _ := s.Recv(0)
		fmt.Printf("Received reply %d [ %s ]\n", i, msg)
	}

	// Output:
	// Connecting to the server...
	// Sending request 0...
	// Received reply 0 [ World ]
	// Sending request 1...
	// Received reply 1 [ World ]
	// Sending request 2...
	// Received reply 2 [ World ]
	// Sending request 3...
	// Received reply 3 [ World ]
	// Sending request 4...
	// Received reply 4 [ World ]
	// Sending request 5...
	// Received reply 5 [ World ]
	// Sending request 6...
	// Received reply 6 [ World ]
	// Sending request 7...
	// Received reply 7 [ World ]
	// Sending request 8...
	// Received reply 8 [ World ]
	// Sending request 9...
	// Received reply 9 [ World ]
}
