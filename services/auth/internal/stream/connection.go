package stream

import (
	"Proyect-Y/auth-service/internal/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection

func getConnection() (*amqp.Connection, error) {
	if connection != nil {
		return connection, nil
	}

	var err error
  env := util.GetEnv()
	connection, err = amqp.Dial(env.AMQP_URI)

	if (err != nil) {
		return nil, nil
	}

	return connection, err
}
