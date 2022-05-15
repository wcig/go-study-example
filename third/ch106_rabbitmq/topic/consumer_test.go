package simple

import (
	"log"
	"testing"

	"github.com/streadway/amqp"
)

// topic工作模式
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
		ExchangeTypeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	go startConsumer1(ch)
	go startConsumer2(ch)

	// 等待退出
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever

	// Output:
	// 2022/05/15 22:11:15  [*] Waiting for messages. To exit press CTRL+C
	// 2022/05/15 22:11:20 >> queue [queue_topic_2] consumer received a message: msg-lazy-2022-05-15 22:11:20
	// 2022/05/15 22:11:20 >> queue [queue_topic_2] consumer received a message: msg-lazy.orange-2022-05-15 22:11:20
	// 2022/05/15 22:11:20 >> queue [queue_topic_1] consumer received a message: msg-lazy.orange.rabbit-2022-05-15 22:11:20
	// 2022/05/15 22:11:20 >> queue [queue_topic_2] consumer received a message: msg-lazy.orange.rabbit-2022-05-15 22:11:20
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

	err = ch.QueueBind(q.Name, routingKey1, exchangeName, false, nil)
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

	err = ch.QueueBind(q.Name, routingKey2, exchangeName, false, nil)
	failOnError(err, "Failed to bind a queue")
	err = ch.QueueBind(q.Name, routingKey3, exchangeName, false, nil)
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
