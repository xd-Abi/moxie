package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/auth"
)

var (
	log = logging.New()
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) GetHealth(ctx context.Context, request *auth.HealthRequest) (*auth.HealthResponse, error) {
	return &auth.HealthResponse{
		Message:   "I am up and running",
		Timestamp: time.Now().Unix(),
	}, nil
}

func main() {
	app := network.NewMicroServiceServer(8000, log)
	auth.RegisterAuthServiceServer(app.InternalServer, &AuthServiceServer{})
	app.Start()
}
