package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/auth"
	"github.com/xd-Abi/moxie/pkg/proto/health"
	"github.com/xd-Abi/moxie/pkg/proto/profile"
)

var (
	authService      auth.AuthServiceClient
	profileService   profile.ProfileServiceClient
	lastHealthReport health.CheckHealthResponse
)

type HealthServiceServer struct {
	health.UnimplementedHealthServiceServer
}

func (s *HealthServiceServer) CheckHealth(ctx context.Context, request *health.CheckHealthRequest) (*health.CheckHealthResponse, error) {
	return &lastHealthReport, nil
}

func checkHealth() {
	startTime := time.Now().Unix()

	// Auth Service health check
	authHealthRecord := health.Record{}
	if authHealthResponse, err := authService.GetHealth(context.Background(), &auth.HealthRequest{}); err == nil {
		authHealthRecord.Message = authHealthResponse.Message
		authHealthRecord.Healthy = true
		authHealthRecord.Timestamp = authHealthResponse.Timestamp
	} else {
		authHealthRecord.Message = err.Error()
		authHealthRecord.Healthy = false
		authHealthRecord.Timestamp = 0
	}

	// Profile Service health check
	profileHealthRecord := health.Record{}
	if profileHealthResponse, err := profileService.GetHealth(context.Background(), &profile.HealthRequest{}); err == nil {
		profileHealthRecord.Message = profileHealthResponse.Message
		profileHealthRecord.Healthy = true
		profileHealthRecord.Timestamp = profileHealthResponse.Timestamp
	} else {
		profileHealthRecord.Message = err.Error()
		profileHealthRecord.Healthy = false
		profileHealthRecord.Timestamp = 0
	}

	lastHealthReport.Checks = &health.Checks{Auth: &authHealthRecord, Profile: &profileHealthRecord}
	lastHealthReport.Timestamp = startTime
}

func main() {
	log := logging.New()
	config.LoadEnvVariables(log)
	authService = auth.NewAuthServiceClient(network.NewGRPCClientConnection(config.GetUint("AUTH_PORT"), log))
	profileService = profile.NewProfileServiceClient(network.NewGRPCClientConnection(config.GetUint("PROFILE_PORT"), log))

	app := network.NewMicroServiceServer(config.GetUint("HEALTH_PORT"), log)
	health.RegisterHealthServiceServer(app.InternalServer, &HealthServiceServer{})
	go app.Start()

	ticker := time.NewTicker(5 * time.Minute)
	checkHealth()

	for range ticker.C {
		checkHealth()
	}
}
