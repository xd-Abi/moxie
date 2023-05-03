package mongodb

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoServer struct {
	internalClient *mongo.Client
}

type MongoCollection struct {
	internalCollection *mongo.Collection
}

func Connect(url string, log *logging.Log) *MongoServer {
	options := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		log.Fatal("Failed to connect to mongo server: %v", err)
	}

	return &MongoServer{
		internalClient: client,
	}
}

func (s *MongoServer) GetCollection(database string, collection string) *MongoCollection {
	col := s.internalClient.Database(database).Collection(collection)
	return &MongoCollection{
		internalCollection: col,
	}
}

func (c *MongoCollection) FindOne(filter bson.D) (bson.M, error) {
	var result bson.M
	err := c.internalCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MongoCollection) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	result, err := c.internalCollection.InsertOne(context.Background(), document)
	if err != nil {
		return nil, err
	}
	return result, nil
}
