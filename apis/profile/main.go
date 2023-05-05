package main

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/profile"
	"github.com/xd-Abi/moxie/pkg/rabbitmq"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type Profile struct {
	UserId  string `bson:"user_id,omitempty"`
	Tag     string `bson:"tag,omitempty"`
	Picture string `bson:"picture"`
}

type MongoConfig struct {
	Uri        string
	Name       string
	Username   string
	Password   string
	Collection string
}

type Config struct {
	RabbitMQUrl string
	Mongo       *MongoConfig
}

type ProfileServiceServer struct {
	Log               *logging.Log
	profileCollection *mongodb.MongoCollection

	profile.UnimplementedProfileServiceServer
}

func NewProfileServiceServer(config *Config, log *logging.Log) *ProfileServiceServer {
	rabbitMQConnection := rabbitmq.NewConnection(config.RabbitMQUrl, log)
	db := mongodb.Connect(config.Mongo.Uri, config.Mongo.Username, config.Mongo.Password, log)
	profileCollection := db.GetCollection(config.Mongo.Name, config.Mongo.Collection)

	server := ProfileServiceServer{
		Log:               log,
		profileCollection: profileCollection,
	}

	rabbitMQConnection.Consume(rabbitmq.ProfileQueueKey, server.handleEvents)
	return &server
}

func (s *ProfileServiceServer) GetProfile(ctx context.Context, request *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {

	if utils.IsEmptyOrWhitespace(request.UserId) {
		return nil, constants.ErrUserIdEmpty
	}

	var p Profile
	err := s.profileCollection.FindOneAndDecode(bson.D{{Key: "user_id", Value: request.UserId}}, &p)
	if err != nil {
		return nil, constants.ErrProfileNotFound
	}

	return &profile.GetProfileResponse{
		UserId:  p.UserId,
		Tag:     p.Tag,
		Picture: p.Picture,
	}, nil
}

func (s *ProfileServiceServer) GetHealth(ctx context.Context, request *profile.HealthRequest) (*profile.HealthResponse, error) {
	return &profile.HealthResponse{
		Message:   "I am up and running",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (s *ProfileServiceServer) handleEvents(event *rabbitmq.Event) error {
	s.Log.Info("Received event: %s", event)

	switch event.Key {
	case rabbitmq.UserSignUpEventKey:
		{
			var payload rabbitmq.UserSignUpEventPayload
			if err := mapstructure.Decode(event.Payload, &payload); err != nil {
				s.Log.Error("Failed to decode event payload")
				break
			}

			if _, err := s.profileCollection.FindOne(bson.D{{Key: "user_id", Value: payload.Id}}); err == nil {
				s.Log.Error("Profile already exists with the user id: %v", payload.Id)
				break
			}

			profile := Profile{
				UserId:  payload.Id,
				Tag:     "Hey there 👋",
				Picture: "",
			}

			_, err := s.profileCollection.InsertOne(profile)
			if err != nil {
				s.Log.Error("Failed to insert profile: %v", err)
			}
		}
		break
	}

	return nil
}

func main() {
	log := logging.New()
	config.LoadEnvVariables(log)
	app := network.NewMicroServiceServer(config.GetUint("PROFILE_PORT"), log)
	config := Config{
		RabbitMQUrl: config.GetString("PROFILE_RABBITMQ_URL"),
		Mongo: &MongoConfig{
			Uri:        config.GetString("PROFILE_DB_URI"),
			Username:   config.GetString("PROFILE_DB_USERNAME"),
			Password:   config.GetString("PROFILE_DB_PASSWORD"),
			Name:       config.GetString("PROFILE_DB_DATABASE"),
			Collection: config.GetString("PROFILE_DB_COLLECTION"),
		},
	}

	profile.RegisterProfileServiceServer(app.InternalServer, NewProfileServiceServer(&config, log))
	app.Start()
}
