package service

import (
	"Proyect-Y/auth-service/internal/storage"
	"Proyect-Y/auth-service/internal/storage/mongo"
	"Proyect-Y/auth-service/internal/storage/redis"
	"Proyect-Y/auth-service/internal/stream"
	"Proyect-Y/typo"
	"sync"

	"github.com/sirupsen/logrus"
)

type dataService struct {
	wg     sync.WaitGroup
	broker *stream.UserProducer
	mongo  storage.UserStore
	cache  storage.UserStore
	logger *logrus.Logger
}

func NewDataService() (*dataService, error) {
	logger := logrus.New()

	broker, err := stream.NewUserProducer()
	if err != nil {
		return nil, err
	}

	mongo, err := mongo.NewMongoAuthClient()
	if err != nil {
		return nil, err
	}

	cache := redis.NewCache()

	return &dataService{
		mongo:  mongo,
		broker: broker,
		cache:  cache,
		logger: logger,
	}, nil
}

func (sv *dataService) manageErr(component string, err error) {
	if err != nil {
		sv.logger.WithError(err).Error(component + " error: ")
	}
}

func (sv *dataService) UserRegister(rg typo.RegisterData) (*typo.AuthData, error) {
	auth := typo.AuthData{
		Email:    rg.Email,
		Password: rg.Password, // TODO encrypt this
		Roles:    typo.User_Rol,
		UserTag:  rg.UserTag,
		BornDate: rg.BornDate,
	}

	res, err := sv.mongo.Save(auth)
	if err != nil || res == nil {
		return res, err
	}

	msg := typo.UserMessage{
		Id:           res.Id,
		RegisterData: rg,
	}
	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Kafka", sv.broker.POST(msg))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		_, er := sv.cache.Save(auth)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *dataService) GetUser(id string) (*typo.AuthData, error) {
	var res *typo.AuthData = nil

	res, err := sv.cache.Get(id)
	if err != nil || res != nil {
		return res, err
	}

	res, err = sv.mongo.Get(id)
	if err != nil || res == nil {
		return res, err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		_, er := sv.cache.Save(*res)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *dataService) GetUserByTag(tag string) (*typo.AuthData, error) {
	var res *typo.AuthData = nil

	res, err := sv.cache.GetByUserTag(tag)
	if err != nil || res != nil {
		return res, err
	}

	res, err = sv.mongo.GetByUserTag(tag)
	if err != nil || res == nil {
		return res, err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()
		_, er := sv.cache.Save(*res)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *dataService) UpdateUser(usr typo.AuthData) (*typo.AuthData, error) {
	res, err := sv.mongo.Edit(usr)
	if err != nil {
		return nil, err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Kafka", sv.broker.Update(*res))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		_, er := sv.cache.Edit(*res)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *dataService) Delete(id string) error {
	err := sv.mongo.Delete(id)

	if err != nil {
		return err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Kafka", sv.broker.Delete(id))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()
		
		sv.manageErr("Cache", sv.cache.Delete(id))
	}()
	return nil
}

func (sv *dataService) CloseAll() {
	sv.wg.Wait()

	sv.manageErr("Kafka", sv.broker.Close())
	sv.manageErr("Mongo", sv.mongo.Close())
	sv.manageErr("Cache", sv.cache.Close())
}
