package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/profile"
)

type ProfileServiceServer struct {
	Log *logging.Log

	profile.UnimplementedProfileServiceServer
}

func NewProfileServiceServer(log *logging.Log) *ProfileServiceServer {
	return &ProfileServiceServer{
		Log: log,
	}
}

func (s *ProfileServiceServer) GetHealth(ctx context.Context, request *profile.HealthRequest) (*profile.HealthResponse, error) {
	return &profile.HealthResponse{
		Message:   "I am up and running",
		Timestamp: time.Now().Unix(),
	}, nil
}

func main() {
	log := logging.New()

	config.LoadEnvVariables(log)
	app := network.NewMicroServiceServer(config.GetUint("PROFILE_PORT"), log)
	profile.RegisterProfileServiceServer(app.InternalServer, NewProfileServiceServer(log))
	app.Start()
}
