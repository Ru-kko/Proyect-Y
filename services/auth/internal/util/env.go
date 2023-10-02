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

	KAFKA_BROKERS string // NOTE ip's separated by "-", example "localhost:8000-172.0.0.1:4000"

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
