package fair

import (
	"fmt"
	"log"
	"testing"

	"github.com/streadway/amqp"
)

// 工作队列模式-循环调度
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
		"queue_work_roundrobin", // 队列名称
		true,                    // 持久化
		false,                   // 自动删除delete when unused
		false,                   // 排他性
		false,                   // 不等待
		nil,                     // 其他参数
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息
	for i := 1; i <= 10; i++ {
		body := fmt.Sprintf("msg-%d", i)
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
		fmt.Printf("send msg %s\n", body)
	}

	fmt.Println(">> producer send message success")

	// Output:
	// send msg msg-1
	// send msg msg-2
	// send msg msg-3
	// send msg msg-4
	// send msg msg-5
	// send msg msg-6
	// send msg msg-7
	// send msg msg-8
	// send msg msg-9
	// send msg msg-10
	// >> producer send message success
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
