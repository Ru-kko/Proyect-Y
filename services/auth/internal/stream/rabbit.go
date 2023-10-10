package stream

import (
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/typo"
	"Proyect-Y/typo/constants/messagging"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserProducer struct {
	channel *amqp.Channel
}

func NewUserProducer() (*UserProducer, error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(messagging.UserTopic, "topic", true, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &UserProducer{
		channel: ch,
	}, nil
}

func (p *UserProducer) produce(data []byte) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := p.channel.PublishWithContext(
		ctx,
		messagging.UserTopic,
		"databases.*",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
  if err != nil {
    return err
  }

	return nil
}

func (p *UserProducer) Close() error {
	return p.channel.Close()
}

func (p *UserProducer) POST(usr typo.UserMessage) error {
	data := typo.ActionPublisher[typo.UserMessage] {
		Data: usr,
		Action: messagging.Save,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = p.produce([]byte(encoded))

	return err
}

func (p *UserProducer) Delete(id string) error {
	data := typo.ActionPublisher[string]{
		Data: id,
		Action: messagging.Delete,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = p.produce([]byte(encoded))

	return err
}

func (p *UserProducer) Update(usr domain.StoredUser) error {
	data := typo.ActionPublisher[typo.AuthChange] {
		Data: typo.AuthChange{
			Id: usr.Id,
			UserTag: usr.UserTag,
			BornDate: usr.BornDate,
		},
		Action: messagging.Update,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = p.produce([]byte(encoded))

	return err
}
