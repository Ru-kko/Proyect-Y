package typo

import "Proyect-Y/typo/constants/messagging"

type ActionPublisher[T interface{}] struct {
	Data   T                      `json:"data"`
	Action messagging.QueueAction `json:"action"`
}

type UserMessage struct {
	Id string `json:"id"`
	RegisterData
}
