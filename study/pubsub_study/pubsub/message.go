package pubsub

type Message struct {
	Topic string
	Body  string
}

func NewMessage(topic string, msg string) *Message {
	return &Message{
		Topic: topic,
		Body:  msg,
	}
}

func (m *Message) GetTopic() string {
	return m.Topic
}

func (m *Message) GetMessageBody() string {
	return m.Body
}
