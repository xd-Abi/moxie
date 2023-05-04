package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/auth"
	"github.com/xd-Abi/moxie/pkg/proto/health"
)

var (
	log              = logging.New()
	authService      auth.AuthServiceClient
	lastHealthReport health.CheckHealthResponse
)

type HealthServiceServer struct {
	health.UnimplementedHealthServiceServer
}

func (s *HealthServiceServer) CheckHealth(ctx context.Context, request *health.CheckHealthRequest) (*health.CheckHealthResponse, error) {
	return &lastHealthReport, nil
}

func checkHealth() {
	authHealthRecord := health.Record{Timestamp: time.Now().Unix()}
	if authHealthResponse, err := authService.GetHealth(context.Background(), &auth.HealthRequest{}); err == nil {
		authHealthRecord.Message, authHealthRecord.Healthy = authHealthResponse.Message, true
	} else {
		authHealthRecord.Message, authHealthRecord.Healthy = err.Error(), false
	}

	lastHealthReport.Checks = &health.Checks{Auth: &authHealthRecord}
	lastHealthReport.Timestamp = authHealthRecord.Timestamp
}

func main() {
	authService = auth.NewAuthServiceClient(network.NewGRPCClientConnection(8000, log))
	app := network.NewMicroServiceServer(8001, log)
	health.RegisterHealthServiceServer(app.InternalServer, &HealthServiceServer{})
	go app.Start()

	ticker := time.NewTicker(5 * time.Minute)
	checkHealth()

	for range ticker.C {
		checkHealth()
	}
}
