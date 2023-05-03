package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/hello"
)

var (
	log = logging.New()
)

type HelloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) SendWelcomeEmail(ctx context.Context, request *hello.WelcomeEmailRequest) (*hello.WelcomeEmailResponse, error) {

	return &hello.WelcomeEmailResponse{
		Success: true,
	}, nil
}

func main() {
	// @TODO: parse port from environment variable
	app := network.NewMicroServiceServer(8000, log)
	hello.RegisterHelloServiceServer(app.InternalServer, &HelloServiceServer{})
	app.Start()
}
