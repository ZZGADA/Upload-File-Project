package config

import "github.com/streadway/amqp"

// 连接信息
const MQURL = "amqp://%s:%s@%d:%d/%d"

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

// 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

func initRabbitMq() {

}

// Publish 发布消息.
//func (ch *Channel) Publish(exchange, key string, body []byte) (err error) {
//	_, err = ch.Channel.PublishWithDeferredConfirmWithContext(context.Background(), exchange, key, false, false,
//		amqp.Publishing{ContentType: "text/plain", Body: body})
//	return err
//}
