package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	reset_password "github.com/xd-Abi/moxie/pkg/proto/reset-password"
)

var (
	log          = logging.New()
	dbCollection *mongodb.MongoCollection
)

type ResetPasswordServiceServer struct {
	reset_password.UnimplementedResetPasswordServiceServer
}

func (s *ResetPasswordServiceServer) ResetPassword(ctx context.Context, request *reset_password.ResetPasswordRequest) (*reset_password.ResetPasswordResponse, error) {
	return nil, nil
}

func main() {
	config.LoadEnvVariables(log)
	db := mongodb.Connect(config.GetString("RESETPASSWORD_DB_HOST"), config.GetString("RESETPASSWORD_DB_USERNAME"), config.GetString("RESETPASSWORD_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("RESETPASSWORD_DB_DATABASE"), config.GetString("RESETPASSWORD_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("RESETPASSWORD_PORT"), log)
	reset_password.RegisterResetPasswordServiceServer(app.InternalServer, &ResetPasswordServiceServer{})
	app.Start()
}
