package application

import (
	"Proyect-Y/auth-service/internal/storage/mongo"
	"Proyect-Y/auth-service/internal/storage/redis"
	"Proyect-Y/typo"

	"github.com/sirupsen/logrus"
)

func UserRegister(rg typo.RegisterData) (*typo.AuthData, error) {
	logger := logrus.New()
	client, err := mongo.NewMongoAuthClient()
	cache := redis.NewCache()

	if err != nil {
		return nil, err
	}

	auth := typo.AuthData{
		Email:    rg.Email,
		Password: rg.Password, // TODO encrypt this
		Roles:    typo.User_Rol,
		UserTag:  rg.UserTag,
		BornDate: rg.BornDate,
	}

	res, err := client.Save(auth)
	if err != nil {
		return nil, err
	}

	cache.Save(auth)

	if err != nil {
		logger.WithError(err).Error("Redis error:")
	}

	return res, nil
}
