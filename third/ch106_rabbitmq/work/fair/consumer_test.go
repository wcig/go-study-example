package fair

import (
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// 工作队列模式-公平调度
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
	// 2022/05/15 18:39:25  [*] Waiting for messages. To exit press CTRL+C
	// 2022/05/15 18:39:26 >> consumer-2 received a message: msg-1
	// 2022/05/15 18:39:26 >> consumer-1 received a message: msg-2
	// 2022/05/15 18:39:27 >> consumer-1 received a message: msg-4
	// 2022/05/15 18:39:28 >> consumer-2 received a message: msg-3
	// 2022/05/15 18:39:28 >> consumer-1 received a message: msg-5
	// 2022/05/15 18:39:29 >> consumer-1 received a message: msg-7
	// 2022/05/15 18:39:30 >> consumer-2 received a message: msg-6
	// 2022/05/15 18:39:30 >> consumer-1 received a message: msg-8
	// 2022/05/15 18:39:31 >> consumer-1 received a message: msg-10
	// 2022/05/15 18:39:32 >> consumer-2 received a message: msg-9
}

func startConsumer(ch *amqp.Channel, id int, sleep int) {
	// 声明队列
	q, err := ch.QueueDeclare(
		"queue_work_fair", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// 从队列消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf(">> consumer-%d received a message: %s", id, d.Body)
			err = d.Ack(false)
			failOnError(err, "Failed to ack message")
			time.Sleep(time.Millisecond * time.Duration(sleep))
		}
	}()
}
