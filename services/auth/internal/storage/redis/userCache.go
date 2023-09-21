package redis

import (
	"context"
	"encoding/json"

	"Proyect-Y/typo"
	"Proyect-Y/typo/constants"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	client *redis.Client
}

func NewCache() *UserCache {
	client := connection()

	return &UserCache{
		client: client,
	}
}

func (cl *UserCache) Get(id string) (*typo.AuthData, error) {
  ctx := context.TODO()

  res, err := cl.client.Get(ctx, id).Bytes()
  
  if err != nil {
    return nil, err
  }

	if res == nil {
		return nil, nil
	}

	usr := &typo.AuthData{}

	if err = json.Unmarshal(res, usr); err != nil {
		return nil, err
	}

	return usr, nil
}

func (cl *UserCache) Save(user typo.AuthData) (*typo.AuthData, error) {
	err := cl.client.Set(context.TODO(), user.Id, user, constants.ExpTime).Err()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (cl *UserCache) Delete(id string) error {
	err := cl.client.Del(context.TODO(), id).Err()
	return err
}

func (cl *UserCache) Edit(usr typo.AuthData) error {
	err := cl.client.Set(context.TODO(), usr.Id, usr, constants.ExpTime).Err()
	return err
}
