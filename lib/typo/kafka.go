package typo

import "Proyect-Y/typo/constants/kafka"

type ActionPublisher[T interface{}] struct {
	Data   T                 `json:"data"`
	Action kafka.KafkaAction `json:"action"`
}

type UserMessage struct {
	Id  string `json:"id"`
	RegisterData
}
