package simple

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// header工作模式
const (
	ExchangeTypeDirect  = "direct"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeTopic   = "topic"
	ExchangeTypeHeaders = "headers"
)

const (
	exchangeName = "exchange_header"
	routingKey   = ""
	queueName1   = "queue_header_1"
	queueName2   = "queue_header_2"
	queueName3   = "queue_header_3"
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
		ExchangeTypeHeaders,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	// 发送消息
	args := amqp.Table{"x": 1}
	sendMsg(ch, args)
	time.Sleep(time.Second)

	args = amqp.Table{"y": 1}
	sendMsg(ch, args)
	time.Sleep(time.Second)

	args = amqp.Table{"x": 1, "y": 1}
	sendMsg(ch, args)
	time.Sleep(time.Second)

	// Output:
	// >> producer send message success: msg-{"x":1}-2022-05-22 17:00:14
	// >> producer send message success: msg-{"y":1}-2022-05-22 17:00:15
	// >> producer send message success: msg-{"x":1,"y":1}-2022-05-22 17:00:16
}

func sendMsg(ch *amqp.Channel, args amqp.Table) {
	nowTimeStr := time.Now().Format("2006-01-02 15:04:05")
	headerStr := toJsonStr(args)
	body := fmt.Sprintf("msg-%s-%s", headerStr, nowTimeStr)
	err := ch.Publish(
		exchangeName, // 交换机名称
		routingKey,   // 路由key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:     args,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(">> producer send message success: %s\n", body)
}

func toJsonStr(args amqp.Table) string {
	data, _ := json.Marshal(args)
	return string(data)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
