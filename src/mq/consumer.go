package mq

import (
	"UploadFileProject/src/entity/dto"
	"UploadFileProject/src/global/enum"
	"UploadFileProject/src/mq/service"
)

// Consumer 消费者订阅
func Consumer() {
	// 初始化mq
	mq := rabbitMqClient
	//defer mq.ReleaseRes() // 完成任务释放资源

	if closed := CheckRabbitClosed(mq.Channel); closed == 1 {
		// channel 已关闭，重连一下
		ReInitChannel()
	}

	// 1.声明队列（两端都要声明，原因在生产者处已经说明）
	_, err := mq.Channel.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		mq.QueueName, // 队列名
		true,         // 是否持久化
		false,        // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,        // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,        // 是否阻塞
		nil,          // 额外属性（我还不会用）
	)
	if err != nil {
		LogMq.Error("声明队列失败", err)
		return
	}

	err = rabbitMqClient.Channel.Qos(
		1,     // prefetch count 服务器将在收到确认之前将那么多消息传递给消费者。
		0,     // prefetch size  服务器将尝试在收到消费者的确认之前至少将那么多字节的交付保持刷新到网络
		false, // 当 global 为 true 时，这些 Qos 设置适用于同一连接上所有通道上的所有现有和未来消费者。当为 false 时，Channel.Qos 设置将应用于此频道上的所有现有和未来消费者
	)
	if err != nil {
		LogMq.Error("rabbitmq设置Qos失败, error: %v", err)
	}

	// 2.从队列获取消息（消费者只关注队列）consume方式会不断的从队列中获取消息
	msgChannel, err := mq.Channel.Consume(
		mq.QueueName,             // 队列名
		ConsumerUploadSingleFile, // 消费者名，用来区分多个消费者，以实现公平分发或均等分发策略
		false,                    // 是否自动应答
		false,                    // 是否排他
		false,                    // 是否接收只同一个连接中的消息，若为true，则只能接收别的conn中发送的消息
		true,                     // 队列消费是否阻塞
		nil,                      // 额外属性
	)
	if err != nil {
		LogMq.Error("获取消息失败", err)
		return
	}

	// 阻塞住
	for {
		select {
		case message := <-msgChannel:
			if closed := CheckRabbitClosed(mq.Channel); closed == 1 {
				ReInitChannel()
			}

			// 消息消费的处理中心
			msgData := message.Body
			obj, messageData, errT := transformMessage(msgData)
			if errT != nil {
				LogMq.Errorf("message 对象转换失败，%#v", errT)
			}

			switch messageData.TaskSituation {
			case enum.TaskSingleFileUpload.ToInt64():
				// 判言直接判断
				fileSingleLoad, _ := obj.(*dto.UpLoadSingleFileOSSMqDTO)
				service.UploadFileOssServiceImpl.UploadSingleFileOSS(fileSingleLoad)
			}

			if errA := message.Ack(true); errA != nil {
				LogMq.Error("mq ack error")
			}
			LogMq.Infof("mq ack success!!,message id is %#v", message.MessageId)
		}

	}
}
