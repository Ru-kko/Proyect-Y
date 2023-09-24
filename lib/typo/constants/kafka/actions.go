package kafka

type KafkaAction string

const (
	KafkaDelete = "delete"
	KafkaUpdate = "update"
	KafkaSave   = "create"
)
