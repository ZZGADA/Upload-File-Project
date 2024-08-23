package mq

import (
	"UploadFileProject/src/entity/dto"
	"encoding/json"
	"fmt"
	"reflect"
)

// TypeRegistry 定义一个全局的类型注册表
var TypeRegistry = make(map[string]reflect.Type)

// 注册类型
func registerType(name string, typ reflect.Type) {
	TypeRegistry[name] = typ
}

// 根据类型名称创建对象
func createObjectByName(name string) (interface{}, error) {
	typ, exists := TypeRegistry[name]
	if !exists {
		return nil, fmt.Errorf("type %s not found", name)
	}
	return reflect.New(typ).Interface(), nil
}

func transformMessage(msgData []byte) (interface{}, *dto.Message, error) {
	var mqMessage = &dto.Message{}
	if errM := json.Unmarshal(msgData, mqMessage); errM != nil {
		LogMq.Errorf("解析rabbitmq数据 %v 失败, error: %v", msgData, errM)
	}

	// 根据类型名称创建对象
	obj, err := createObjectByName(mqMessage.BodyStructureName)
	if err != nil {
		LogMq.Errorf("Error creating object: %#v", err)
	}

	// 反序列化 JSON 数据到创建的对象
	err = json.Unmarshal([]byte(mqMessage.Message), obj)
	if err != nil {
		LogMq.Errorf("Error unmarshalling JSON: %#v", err)
	}

	return obj, mqMessage, err
}
