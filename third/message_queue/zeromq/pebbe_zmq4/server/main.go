package main

import (
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()
	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	for {
		// Wait for next request from client
		msg, _ := s.Recv(0)
		log.Printf("Received %s\n", msg)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		s.Send("World", 0)
	}

	// Output:
	// 2022/08/04 11:13:22 Received Hello
	// 2022/08/04 11:13:23 Received Hello
	// 2022/08/04 11:13:24 Received Hello
	// 2022/08/04 11:13:25 Received Hello
	// 2022/08/04 11:13:26 Received Hello
	// 2022/08/04 11:13:27 Received Hello
	// 2022/08/04 11:13:28 Received Hello
	// 2022/08/04 11:13:29 Received Hello
	// 2022/08/04 11:13:30 Received Hello
	// 2022/08/04 11:13:31 Received Hello
}
