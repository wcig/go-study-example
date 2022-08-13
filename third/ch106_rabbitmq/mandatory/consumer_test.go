package simple

import (
	"log"
	"testing"

	"github.com/streadway/amqp"
)

// direct工作模式
func TestConsumer(t *testing.T) {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换器
	err = ch.ExchangeDeclare(
		exchangeName,
		ExchangeTypeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	go startConsumer1(ch)
	// go startConsumer2(ch)

	// 等待退出
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever

	// Output:
	// 2022/05/15 21:05:57  [*] Waiting for messages. To exit press CTRL+C
	// 2022/05/15 21:06:00 >> queue [queue_direct_2] consumer received a message: msg-routingkey_direct_info-2022-05-15 21:06:00
	// 2022/05/15 21:06:00 >> queue [queue_direct_2] consumer received a message: msg-routingkey_direct_warning-2022-05-15 21:06:00
	// 2022/05/15 21:06:00 >> queue [queue_direct_2] consumer received a message: msg-routingkey_direct_error-2022-05-15 21:06:00
	// 2022/05/15 21:06:00 >> queue [queue_direct_1] consumer received a message: msg-routingkey_direct_error-2022-05-15 21:06:00
}

func startConsumer1(ch *amqp.Channel) {
	// 声明队列
	q, err := ch.QueueDeclare(
		"queueName1", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "routingKeyError", exchangeName, false, nil)
	failOnError(err, "Failed to bind a queue")

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
			log.Printf(">> queue [%s] consumer received a message: %s", q.Name, d.Body)
		}
	}()
}
