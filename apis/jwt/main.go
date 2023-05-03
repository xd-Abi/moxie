package main

import (
	"context"
	"time"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/jwt"
)

var (
	log = logging.New()
)

type JwtServiceServer struct {
	Secret          string
	TokenExpiration time.Duration
	jwt.UnimplementedJwtServiceServer
}

func (s *JwtServiceServer) GenerateToken(ctx context.Context, request *jwt.GenerateTokenRequest) (*jwt.GenerateTokenResponse, error) {

}

func (s *JwtServiceServer) VerifyToken(ctx context.Context, request *jwt.VerifyTokenRequest) (*jwt.VerifyTokenResponse, error) {

}

func main() {
	config.LoadEnvVariables(log)

	app := network.NewMicroServiceServer(config.GetUint("JWT_PORT"), log)
	jwt.RegisterJwtServiceServer(app.InternalServer, &JwtServiceServer{
		Secret:          config.GetString("JWT_SECRET"),
		TokenExpiration: time.Duration(config.GetUint("JWT_EXPIRATION")),
	})
	app.Start()
}
