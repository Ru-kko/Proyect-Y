package service

import (
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/auth-service/internal/security"
	"Proyect-Y/auth-service/internal/storage"
	"Proyect-Y/auth-service/internal/storage/mongo"
	"Proyect-Y/auth-service/internal/storage/redis"
	"Proyect-Y/auth-service/internal/stream"
	"Proyect-Y/typo"
	"sync"

	"github.com/sirupsen/logrus"
)

type DataService struct {
	wg     sync.WaitGroup
	broker *stream.UserProducer
	mongo  storage.UserStore
	cache  storage.UserStore
	logger *logrus.Logger
}

func NewDataService() (*DataService, error) {
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

	return &DataService{
		mongo:  mongo,
		broker: broker,
		cache:  cache,
		logger: logger,
	}, nil
}

func (sv *DataService) manageErr(component string, err error) {
	if err != nil {
		sv.logger.WithError(err).Error(component + " error: ")
	}
}

func (sv *DataService) UserRegister(rg typo.RegisterData) (*domain.StoredUser, error) {
	hashedPsw, err := security.EncryptPassword(rg.Password)
	if err != nil {
		return nil, err
	}

	auth := domain.StoredUser{
		Email:    rg.Email,
		Password: hashedPsw,
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

		sv.manageErr("Rabbit", sv.broker.POST(msg))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		_, er := sv.cache.Save(auth)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *DataService) GetUser(id string) (*domain.StoredUser, error) {
	var res *domain.StoredUser = nil

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

func (sv *DataService) GetUserByTag(tag string) (*domain.StoredUser, error) {
	var res *domain.StoredUser = nil

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

func (sv *DataService) UpdateUser(usr domain.StoredUser) (*domain.StoredUser, error) {
	res, err := sv.mongo.Edit(usr)
	if err != nil {
		return nil, err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Rabbit", sv.broker.Update(*res))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		_, er := sv.cache.Edit(*res)
		sv.manageErr("Cache", er)
	}()

	return res, nil
}

func (sv *DataService) Delete(id string) error {
	err := sv.mongo.Delete(id)

	if err != nil {
		return err
	}

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Rabbit", sv.broker.Delete(id))
	}()

	go func() {
		sv.wg.Add(1)
		defer sv.wg.Done()

		sv.manageErr("Cache", sv.cache.Delete(id))
	}()
	return nil
}

func (sv *DataService) CloseAll() {
	sv.wg.Wait()

	sv.manageErr("Rabbit", sv.broker.Close())
	sv.manageErr("Mongo", sv.mongo.Close())
	sv.manageErr("Cache", sv.cache.Close())
}
