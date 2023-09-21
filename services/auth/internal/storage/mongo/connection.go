package mongo

import (
	"context"
	"fmt"

	"Proyect-Y/auth-service/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection() (*mongo.Client, error) {
	env := util.GetEnv()
	url := fmt.Sprintf("mongodb://%s:%s@%s/auth", env.MONGO_USR, env.MONGO_PSW, env.MONGO_HOST)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return client, nil
}
