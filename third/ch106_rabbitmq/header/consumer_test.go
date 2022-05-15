package simple

import (
	"log"
	"testing"

	"github.com/streadway/amqp"
)

// header工作模式
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
		ExchangeTypeHeaders,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	go startConsumer1(ch)
	go startConsumer2(ch)
	go startConsumer3(ch)

	// 等待退出
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever

	// Output:
	// 2022/05/15 22:24:25  [*] Waiting for messages. To exit press CTRL+C
	// 2022/05/15 22:24:29 >> queue [queue_header_1] consumer received a message: msg-{"x":1}-2022-05-15 22:24:29
	// 2022/05/15 22:24:29 >> queue [queue_header_3] consumer received a message: msg-{"x":1,"y":1}-2022-05-15 22:24:29
	// 2022/05/15 22:24:29 >> queue [queue_header_2] consumer received a message: msg-{"y":1}-2022-05-15 22:24:29
	// 2022/05/15 22:24:29 >> queue [queue_header_2] consumer received a message: msg-{"x":1,"y":1}-2022-05-15 22:24:29
	// 2022/05/15 22:24:29 >> queue [queue_header_1] consumer received a message: msg-{"x":1,"y":1}-2022-05-15 22:24:29
}

func startConsumer1(ch *amqp.Channel) {
	// 声明队列
	q, err := ch.QueueDeclare(
		queueName1, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	args := amqp.Table{"x": 1}
	err = ch.QueueBind(q.Name, routingKey, exchangeName, false, args)
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

func startConsumer2(ch *amqp.Channel) {
	// 声明队列
	q, err := ch.QueueDeclare(
		queueName2, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	args := amqp.Table{"y": 1}
	err = ch.QueueBind(q.Name, routingKey, exchangeName, false, args)
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

func startConsumer3(ch *amqp.Channel) {
	// 声明队列
	q, err := ch.QueueDeclare(
		queueName3, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	args := amqp.Table{"x": 1, "y": 1}
	err = ch.QueueBind(q.Name, routingKey, exchangeName, false, args)
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
