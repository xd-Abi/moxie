package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	change_password "github.com/xd-Abi/moxie/pkg/proto/change-password"
)

var (
	log          = logging.New()
	dbCollection *mongodb.MongoCollection
)

type ChangePasswordServiceServer struct {
	change_password.UnimplementedChangePasswordServiceServer
}

func (s *ChangePasswordServiceServer) ChangePassword(ctx context.Context, request *change_password.ChangePasswordRequest) (*change_password.ChangePasswordResponse, error) {
	return &change_password.ChangePasswordResponse{
		PasswordChanged: true,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)

	db := mongodb.Connect(config.GetString("CHANGEPASSWORD_DB_HOST"), config.GetString("CHANGEPASSWORD_DB_USERNAME"), config.GetString("CHANGEPASSWORD_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("CHANGEPASSWORD_DB_DATABASE"), config.GetString("CHANGEPASSWORD_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("CHANGEPASSWORD_PORT"), log)
	change_password.RegisterChangePasswordServiceServer(app.InternalServer, &ChangePasswordServiceServer{})
	app.Start()
}
