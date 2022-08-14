package pubsub

import (
	"fmt"
	"sync"
)

type Subscribers map[string]*Subscriber

type Broker struct {
	Subscribers Subscribers
	Topics      map[string]Subscribers
	mu          sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		Subscribers: Subscribers{},
		Topics:      make(map[string]Subscribers),
	}
}

func (b *Broker) AddSubscriber() *Subscriber {
	b.mu.Lock()
	defer b.mu.Unlock()

	id, sub := CreateNewSubscriber()
	b.Subscribers[id] = sub
	return sub
}

func (b *Broker) RemoveSubscriber(s *Subscriber) {
	for topic := range s.Topics {
		b.Unsubscribe(s, topic)
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.Subscribers, s.Id)
	s.Destruct()
}

func (b *Broker) GetSubscribers(topic string) int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.Topics[topic])
}

func (b *Broker) Broadcast(topics []string, msg string) {
	for _, topic := range topics {
		for _, sub := range b.Topics[topic] {
			m := NewMessage(topic, msg)
			go func(s *Subscriber) {
				s.Signal(m)
			}(sub)
		}
	}
}

func (b *Broker) Subscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Topics[topic] == nil {
		b.Topics[topic] = Subscribers{}
	}
	s.AddTopic(topic)
	b.Topics[topic][s.Id] = s
	fmt.Printf("%s Subscribed for topic: %s\n", s.Id, topic)
}

func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.Topics[topic], s.Id)
	s.RemoveTopic(topic)
	fmt.Printf("%s Unsubscribed for topic: %s\n", s.Id, topic)
}

func (b *Broker) Publish(topic string, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	subs := b.Topics[topic]
	for _, sub := range subs {
		m := NewMessage(topic, msg)
		if !sub.Active {
			return
		}
		go func(s *Subscriber) {
			s.Signal(m)
		}(sub)
	}
}
