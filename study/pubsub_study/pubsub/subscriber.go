package pubsub

import (
	"fmt"
	"math/rand"
	"sync"
)

type Subscriber struct {
	Id       string
	Messages chan *Message
	Topics   map[string]bool
	Active   bool
	mu       sync.Mutex
}

func CreateNewSubscriber() (id string, s *Subscriber) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	id = fmt.Sprintf("%X-%X", b[0:4], b[4:8])
	s = &Subscriber{
		Id:       id,
		Messages: make(chan *Message),
		Topics:   make(map[string]bool),
		Active:   true,
	}
	return id, s
}

func (s *Subscriber) AddTopic(topic string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Topics[topic] = true
}

func (s *Subscriber) RemoveTopic(topic string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Topics, topic)
}

func (s *Subscriber) GetTopics() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var topics []string
	for topic := range s.Topics {
		topics = append(topics, topic)
	}
	return topics
}

func (s *Subscriber) Destruct() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Active = false
	close(s.Messages)
}

func (s *Subscriber) Signal(msg *Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Active {
		s.Messages <- msg
	}
}

func (s *Subscriber) Listen() {
	for {
		if msg, ok := <-s.Messages; ok {
			fmt.Printf("Subscriber %s, received: %s from topic: %s\n", s.Id, msg.GetMessageBody(), msg.GetTopic())
		}
	}
}
