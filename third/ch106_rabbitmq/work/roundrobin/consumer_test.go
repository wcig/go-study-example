package fair

import (
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// 工作队列模式-循环调度
func TestConsumer(t *testing.T) {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	go startConsumer(ch, 1, 1000)
	go startConsumer(ch, 2, 2000)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever

	// Output:
	// 2022/05/15 18:01:23 >> consumer-1 received a message: msg-2
	// 2022/05/15 18:01:23 >> consumer-2 received a message: msg-1
	// 2022/05/15 18:01:24 >> consumer-1 received a message: msg-4
	// 2022/05/15 18:01:25 >> consumer-2 received a message: msg-3
	// 2022/05/15 18:01:25 >> consumer-1 received a message: msg-6
	// 2022/05/15 18:01:26 >> consumer-1 received a message: msg-8
	// 2022/05/15 18:01:27 >> consumer-2 received a message: msg-5
	// 2022/05/15 18:01:27 >> consumer-1 received a message: msg-10
	// 2022/05/15 18:01:29 >> consumer-2 received a message: msg-7
	// 2022/05/15 18:01:31 >> consumer-2 received a message: msg-9
}

func startConsumer(ch *amqp.Channel, id int, sleep int) {
	// 声明队列
	q, err := ch.QueueDeclare(
		"queue_work_roundrobin", // name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 从队列消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf(">> consumer-%d received a message: %s", id, d.Body)
			time.Sleep(time.Millisecond * time.Duration(sleep))
		}
	}()
}
