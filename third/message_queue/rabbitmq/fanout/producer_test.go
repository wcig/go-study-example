package simple

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// fanout发布订阅工作模式
const (
	ExchangeTypeDirect  = "direct"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeTopic   = "topic"
	ExchangeTypeHeaders = "headers"
)

const (
	exchangeName = "exchange_fanout"
	routingKey   = "routingkey_fanout"
	queueName1   = "queue_fanout_1"
	queueName2   = "queue_fanout_2"
	queueName3   = "queue_fanout_3"
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
		ExchangeTypeFanout,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	// 发送消息
	body := "msg-" + time.Now().Format("2006-01-02 15:04:05")
	err = ch.Publish(
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

	// Output:
	// >> producer send message success: msg-2022-05-15 20:23:55
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
