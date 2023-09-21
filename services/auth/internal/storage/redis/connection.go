package redis

import (
	"Proyect-Y/auth-service/internal/util"

	"github.com/redis/go-redis/v9"
)

func connection() *redis.Client {
	env := util.GetEnv()
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_HOST,
		Password: env.REDIS_PSW,
		DB:       0,
	})

	return client
}
