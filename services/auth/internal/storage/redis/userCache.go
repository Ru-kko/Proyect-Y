package redis

import (
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/typo/constants"
	"context"
	"encoding/json"
	"sync"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	client *redis.Client
	wg     sync.WaitGroup
}

func NewCache() *UserCache {
	client := connection()

	return &UserCache{
		client: client,
	}
}

func (cl *UserCache) Close() error {
	cl.wg.Wait()
	return cl.client.Close()
}

func (cl *UserCache) GetByUserTag(tag string) (*domain.StoredUser, error) {
	id, err := cl.client.Get(context.TODO(), "tag:"+tag).Result()
	if err != nil {
		switch err {
		case redis.Nil:
			return nil, nil
		default:
			return nil, err
		}
	}

	return cl.Get(id)
}

func (cl *UserCache) Get(id string) (*domain.StoredUser, error) {
	ctx := context.TODO()

	res, err := cl.client.Get(ctx, id).Bytes()
	if err != nil {
		switch err {
		case redis.Nil:
			return nil, nil
		default:
			return nil, err
		}
	}

	go func() {
		cl.wg.Add(1)
		defer cl.wg.Done()
		cl.client.ExpireNX(ctx, id, constants.ExpTime)
	}()

	usr := &domain.StoredUser{}

	if err = json.Unmarshal(res, usr); err != nil {
		return nil, err
	}

	return usr, nil
}

func (cl *UserCache) Save(user domain.StoredUser) (*domain.StoredUser, error) {
	str, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	if err = cl.client.Set(context.TODO(), user.Id, string(str), constants.ExpTime).Err(); err != nil {
		return nil, err
	}

	err = cl.client.Set(context.TODO(), "tag:"+user.UserTag, user.Id, constants.ExpTime).Err()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (cl *UserCache) Delete(id string) error {
	err := cl.client.Del(context.TODO(), id).Err()
	return err
}

func (cl *UserCache) Edit(usr domain.StoredUser) (*domain.StoredUser, error) {
	return cl.Save(usr)
}
