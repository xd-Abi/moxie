package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/profile"
	"github.com/xd-Abi/moxie/pkg/rabbitmq"
)

type Config struct {
	RabbitMQUrl string
}

type ProfileServiceServer struct {
	Log *logging.Log

	profile.UnimplementedProfileServiceServer
}

func NewProfileServiceServer(config *Config, log *logging.Log) *ProfileServiceServer {
	rabbitMQConnection := rabbitmq.NewConnection(config.RabbitMQUrl, log)

	server := ProfileServiceServer{
		Log: log,
	}

	rabbitMQConnection.Consume(rabbitmq.ProfileQueueKey, server.handleEvents)
	return &server
}

func (s *ProfileServiceServer) GetHealth(ctx context.Context, request *profile.HealthRequest) (*profile.HealthResponse, error) {
	return &profile.HealthResponse{
		Message:   "I am up and running",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (s *ProfileServiceServer) handleEvents(event *rabbitmq.Event) error {

	s.Log.Info("Received event: %s", event)

	return nil
}

func main() {
	log := logging.New()
	config.LoadEnvVariables(log)
	app := network.NewMicroServiceServer(config.GetUint("PROFILE_PORT"), log)
	config := Config{
		RabbitMQUrl: config.GetString("PROFILE_RABBITMQ_URL"),
	}

	profile.RegisterProfileServiceServer(app.InternalServer, NewProfileServiceServer(&config, log))
	app.Start()
}
