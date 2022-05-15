package simple

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// topic工作模式
const (
	ExchangeTypeDirect  = "direct"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeTopic   = "topic"
	ExchangeTypeHeaders = "headers"
)

const (
	exchangeName = "exchange_topic"
	routingKey1  = "*.orange.*"
	routingKey2  = "*.*.rabbit"
	routingKey3  = "lazy.#"
	queueName1   = "queue_topic_1"
	queueName2   = "queue_topic_2"
)

func TestProducer(t *testing.T) {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//  创建通道
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

	// 发送消息
	routingKey := "lazy"
	sendMsg(ch, routingKey)

	routingKey = "lazy.orange"
	sendMsg(ch, routingKey)

	routingKey = "lazy.orange.rabbit"
	sendMsg(ch, routingKey)

	// Output:
	// >> producer send message success: msg-lazy-2022-05-15 22:11:20
	// >> producer send message success: msg-lazy.orange-2022-05-15 22:11:20
	// >> producer send message success: msg-lazy.orange.rabbit-2022-05-15 22:11:20
}

func sendMsg(ch *amqp.Channel, routingKey string) {
	nowTimeStr := time.Now().Format("2006-01-02 15:04:05")
	body := fmt.Sprintf("msg-%s-%s", routingKey, nowTimeStr)
	err := ch.Publish(
		exchangeName, // 交换机名称
		routingKey,   // 路由key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(">> producer send message success: %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
