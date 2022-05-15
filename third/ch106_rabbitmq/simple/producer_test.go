package simple

import (
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
)

// 简单队列工作模式
func TestProducer(t *testing.T) {
	// 创建连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//  创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"queue_simple", // 队列名称
		false,          // 持久化
		false,          // 自动删除delete when unused
		false,          // 排他性
		false,          // 不等待
		nil,            // 其他参数
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 路由key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println(">> producer send message success")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
