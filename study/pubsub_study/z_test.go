package main

import (
	"fmt"
	"go-app/study/pubsub_study/pubsub"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	broker := pubsub.NewBroker()
	sub1 := broker.AddSubscriber()
	topic1 := "test1"
	sub1.AddTopic(topic1)
	broker.Subscribe(sub1, topic1)
	go func() {
		i := 0
		for {
			i++
			broker.Publish(topic1, fmt.Sprintf("msg-%d", i))
			time.Sleep(time.Second)
		}
	}()
	sub1.Listen()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
}
