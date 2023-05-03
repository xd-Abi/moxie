package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/signup"
)

var (
	log = logging.New()
)

type SignUpServiceServer struct {
	signup.UnimplementedSignUpServiceServer
}

func (s *SignUpServiceServer) SignUp(ctx context.Context, request *signup.SignUpRequest) (*signup.SignUpResponse, error) {

	return nil, nil
}

func main() {
	config.LoadEnvVariables(log)

	app := network.NewMicroServiceServer(config.GetUint("SIGN_UP_PORT"), log)
	signup.RegisterSignUpServiceServer(app.InternalServer, &SignUpServiceServer{})
	app.Start()
}
