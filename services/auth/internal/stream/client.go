package stream

import (
	"Proyect-Y/auth-service/internal/util"
	"Proyect-Y/typo"
	"Proyect-Y/typo/constants/kafka"
	"encoding/json"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type UserProducer struct {
	producer sarama.SyncProducer
}

func NewUserProducer() (*UserProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	hosts := strings.Split(util.GetEnv().KAFKA_BROKERS, "-")

	producer, err := sarama.NewSyncProducer(hosts, config)
	if err != nil {
		return nil, err
	}

	return &UserProducer{
		producer,
	}, nil
}

func (p *UserProducer) Close() error {
	return p.producer.Close()
}

func (p *UserProducer) POST(usr typo.UserMessage) error {
	data := typo.ActionPublisher[typo.UserMessage]{
		Data:   usr,
		Action: kafka.KafkaSave,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic:     kafka.UserTopic,
		Partition: -1,
		Value:     sarama.ByteEncoder(encoded),
		Timestamp: time.Now(),
	}

	_, _, err = p.producer.SendMessage(message)

	return err
}

func (p *UserProducer) Delete(id string) error {
	data := typo.ActionPublisher[string]{
		Data:   id,
		Action: kafka.KafkaSave,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic:     kafka.KafkaDelete,
		Partition: -1,
		Value:     sarama.ByteEncoder(encoded),
		Timestamp: time.Now(),
	}

	_, _, err = p.producer.SendMessage(message)

	return err
}

func (p *UserProducer) Update(changes typo.AuthData) error {
	data := typo.ActionPublisher[typo.AuthChange]{
		Data: typo.AuthChange{
			Id: changes.Id,
			UserTag: changes.UserTag,
			BornDate: changes.BornDate,
		},
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: kafka.KafkaUpdate,
		Partition: -1,
		Value: sarama.ByteEncoder(encoded),
		Timestamp: time.Now(),
	}

	_, _, err = p.producer.SendMessage(message)

	return err
}
