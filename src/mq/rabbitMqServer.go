package mq

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global"
	"UploadFileProject/src/mq/service"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"reflect"
)

var RabbitMqConfig *global.RabbitMqConfig
var rabbitMqClient *RabbitMQ
var LogMq *logrus.Logger

// UrlOrigin //账号：密码@rabbitmq服务器地址：端口号/vhost
const UrlOrigin = "amqp://%s:%s@%s:%d/%s"
const ConsumerUploadSingleFile = "ConsumerUploadSingleFile"
const ContentType = "application/json"

var Url string

// RabbitMQ rabbitMQ结构体
type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	RoutingKey string
	//连接信息
	Url         string
	ContentType string
}

func InitRabbitMqServer(mqConfig *global.RabbitMqConfig) {
	RabbitMqConfig = mqConfig
	Url = fmt.Sprintf(UrlOrigin,
		RabbitMqConfig.Username,
		RabbitMqConfig.Password,
		RabbitMqConfig.Host,
		RabbitMqConfig.Port,
		RabbitMqConfig.VirtualHost)
	rabbitMqClient = NewRabbitMq()

	registerType("UpLoadSingleFileOSSMqDTO", reflect.TypeOf(dto.UpLoadSingleFileOSSMqDTO{}))

	LogMq = global.Log

	service.InitConsumerService(global.Log)
}

// NewRabbitMq 创建结构体实例
func NewRabbitMq() *RabbitMQ {
	rabbitMQ := &RabbitMQ{
		QueueName:   RabbitMqConfig.ServerOne.Queue,
		Exchange:    RabbitMqConfig.ServerOne.Exchange,
		RoutingKey:  RabbitMqConfig.ServerOne.RoutingKey,
		Url:         Url,
		ContentType: ContentType,
	}

	var err error
	//创建rabbitmq连接
	rabbitMQ.Connection, err = amqp.Dial(rabbitMQ.Url)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Connection.Channel()
	checkErr(err, "创建channel失败")

	return rabbitMQ
}

// 错误处理
func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}

// ReleaseRes 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	if err := mq.Connection.Close(); err != nil {
		global.Log.Error("mq.Connection.Close failed")
	}
	if err := mq.Channel.Close(); err != nil {
		global.Log.Error("mq.channel.Close failed")
	}
}

// CheckRabbitClosed 0表示channel未关闭，1表示channel已关闭
func CheckRabbitClosed(ch *amqp.Channel) int64 {
	d := reflect.ValueOf(*ch)
	i := d.FieldByName("closed").Int()
	return i
}

// ReInitChannel channel重连接
func ReInitChannel() {
	rabbitMqClient = NewRabbitMq()
	err := rabbitMqClient.Channel.Qos(1, 0, false)
	if err != nil {
		LogMq.Error("rabbitmq重连后设置Qos失败, error: %v", err)
	}
}

func NewMessage(data interface{}, dataStructureName string, taskId int64) *dto.Message {
	val, _ := json.Marshal(data)
	return &dto.Message{
		Message:           string(val),
		BodyStructureName: dataStructureName,
		TaskSituation:     taskId,
	}
}
