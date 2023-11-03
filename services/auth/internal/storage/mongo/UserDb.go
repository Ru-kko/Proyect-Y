package mongo

import (
	"Proyect-Y/auth-service/internal/domain"
	"Proyect-Y/common/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	coll *mongo.Collection
	conn *mongo.Client
}

func NewMongoAuthClient() (*MongoClient, error) {
	client, err := connection()
	if err != nil {
		return nil, err
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

	return &MongoClient{
		coll: collection,
		conn: client,
	}, nil
}

func (cl *MongoClient) Close() error {
	return cl.conn.Disconnect(context.TODO())
}

func (cl *MongoClient) GetByUserTag(tag string) (*domain.StoredUser, error) {
	res := &domain.StoredUser{}
	data := cl.coll.FindOne(context.TODO(), bson.M{"user_tag": tag})

	err := data.Decode(res)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, nil
		default:
			return nil, err
		}
	}

	return res, err
}

func (cl *MongoClient) Get(id string) (*domain.StoredUser, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	data := cl.coll.FindOne(context.TODO(), bson.M{"_id": objid})

	res := &domain.StoredUser{}
	err = data.Decode(res)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, nil
		default:
			return nil, err
		}
	}

	return res, nil
}

func (cl *MongoClient) Save(usr domain.StoredUser) (*domain.StoredUser, error) {
	data, err := cl.coll.InsertOne(context.TODO(), usr)
	if err != nil {
		return nil, err
	}

	usr.Id = data.InsertedID.(primitive.ObjectID).Hex()

	return &usr, nil
}

func (cl *MongoClient) Edit(user domain.StoredUser) (*domain.StoredUser, error) {
	clone := user
	clone.Id = ""
	parsed := util.StructToMap(clone)

	_, err := cl.coll.UpdateOne(context.TODO(), bson.M{"_id": user.Id}, parsed)
	if err != nil {
		return nil, err
	}

	return cl.Get(user.Id)
}

func (cl *MongoClient) Delete(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = cl.coll.DeleteOne(context.TODO(), bson.M{"_id": objId})

	return err
}
