package util

import (
	ev "Proyect-Y/common/env"
)

type enviroment struct {
	PORT int

	MONGO_HOST string
	MONGO_PSW  string
	MONGO_USR  string

	REDIS_HOST string
	REDIS_PSW  string

	AMQP_URI string

	JWT_SECRET string
}

var env *enviroment

func GetEnv() enviroment {
	if env != nil {
		return *env
	}

	env = &enviroment{}

	ev.Load()
	ev.Parse[enviroment](env)

	return *env
}
