package dto

// Message Mq消息对象
type Message struct {
	Message           string `json:"message" mapstructure:"message"`
	BodyStructureName string `json:"bodyStructureName" mapstructure:"bodyStructureName"`
	TaskSituation     int64  `json:"taskSituation" mapstructure:"taskSituation"`
}
