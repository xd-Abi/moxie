package mongodb

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoServer struct {
	internalClient *mongo.Client
}

type MongoCollection struct {
	internalCollection *mongo.Collection
}

func Connect(url string, username string, password string, log *logging.Log) *MongoServer {
	opt := options.Client().ApplyURI(url)
	opt.Auth = &options.Credential{
		Username: username,
		Password: password,
	}
	client, err := mongo.Connect(context.Background(), opt)

	if err != nil {
		log.Fatal("Failed to connect to mongo server: %v", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Failed to connect to mongo server: %v", err)
	}

	log.Info("Connected to database")

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
