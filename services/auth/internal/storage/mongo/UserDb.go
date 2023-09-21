package mongo

import (
	"Proyect-Y/auth-service/internal/util"
	"Proyect-Y/typo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	coll *mongo.Collection
}

func NewMongoAuthClient() (MongoClient, error) {
	client, err := connection()

	if err != nil {
		return MongoClient{}, err
	}

	collection := client.Database("auth").Collection("users")
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_tag", Value: -1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: -1}},
			Options: options.Index().SetUnique(true),
		},
	}
	collection.Indexes().CreateMany(context.TODO(), indexes)

	return MongoClient{
		coll: collection,
	}, nil
}

func (cl *MongoClient) Get(id string) (*typo.AuthData, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	data := cl.coll.FindOne(context.TODO(), bson.M{"_id": objid})

	res := typo.AuthData{}
	err = data.Decode(res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (cl *MongoClient) Save(usr typo.AuthData) (*typo.AuthData, error) {
	data, err := cl.coll.InsertOne(context.TODO(), usr)
	if err != nil {
		return nil, err
	}

	usr.Id = data.InsertedID.(string)

	return &usr, nil
}

func (cl *MongoClient) Edit(user typo.AuthData) error {
	clone := user
	clone.Id = ""
	parsed := util.StructToMap(clone)

	_, err := cl.coll.UpdateOne(context.TODO(), bson.M{"_id": user.Id}, parsed)

	return err
}

func (cl *MongoClient) Delete(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = cl.coll.DeleteOne(context.TODO(), bson.M{"_id": objId})

	return err
}
