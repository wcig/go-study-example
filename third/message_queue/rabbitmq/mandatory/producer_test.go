package simple

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

// direct工作模式
const (
	ExchangeTypeDirect  = "direct"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeTopic   = "topic"
	ExchangeTypeHeaders = "headers"
)

const (
	exchangeName = "exchange_direct_mandatory_true"
	// routingKeyInfo    = "routingkey_direct_info"
	// routingKeyWarning = "routingkey_direct_warning"
	// routingKeyError   = "routingkey_direct_error"
	// queueName1        = "queue_direct_1"
	// queueName2        = "queue_direct_2"
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
		ExchangeTypeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	// 发送消息
	// routingKeys := []string{routingKeyInfo, routingKeyWarning, routingKeyError}
	// routingKeys := []string{""}
	// for _, routingKey := range routingKeys {

	var ack = make(chan uint64)
	var nack = make(chan uint64)
	ch.NotifyConfirm(ack, nack)

	returnChannel := make(chan amqp.Return)
	ch.NotifyReturn(returnChannel)

	routingKey := ""
	nowTimeStr := time.Now().Format("2006-01-02 15:04:05")
	body := fmt.Sprintf("msg-%s-%s", routingKey, nowTimeStr)
	err = ch.Publish(
		exchangeName, // 交换机名称
		routingKey,   // 路由key
		true,         // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(">> producer send message success: %s\n", body)

	select {
	case returnNotification := <-returnChannel:
		fmt.Println(">> return channel")
		printJson(returnNotification)
	case <-ack:
		fmt.Println(">> ack")
	case <-nack:
		fmt.Println(">> nack")
	}
	// }

	// Output:
	// >> producer send message success: msg-routingkey_direct_info-2022-05-15 21:06:00
	// >> producer send message success: msg-routingkey_direct_warning-2022-05-15 21:06:00
	// >> producer send message success: msg-routingkey_direct_error-2022-05-15 21:06:00
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func printJson(i interface{}) {
	data, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(data))
}
