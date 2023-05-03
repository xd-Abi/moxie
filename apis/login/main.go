package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/jwt"
	"github.com/xd-Abi/moxie/pkg/proto/login"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	log          = logging.New()
	jwtService   jwt.JwtServiceClient
	dbCollection *mongodb.MongoCollection
)

type LoginServiceServer struct {
	login.UnimplementedLoginServiceServer
}

func (s *LoginServiceServer) Login(ctx context.Context, request *login.LoginRequest) (*login.LoginResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Email) {
		return nil, constants.ErrEmailEmpty
	}
	if !utils.IsEmail(request.Email) {
		return nil, constants.ErrEmailInvalid
	}
	if utils.IsEmptyOrWhitespace(request.Password) {
		return nil, constants.ErrPasswordEmpty
	}

	user, err := dbCollection.FindOne(bson.D{{Key: "email", Value: request.Email}})
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	if currentHashedPassword, ok := user["password"].(string); ok {
		if !utils.ComparePasswords([]byte(currentHashedPassword), []byte(request.Password)) {
			return nil, constants.ErrPasswordInvalid
		}
	} else {
		log.Error("Failed to convert current password into string")
		return nil, constants.ErrInternal
	}

	userId, ok := user["id"].(string)
	if !ok {
		log.Error("Failed to convert user id into string")
		return nil, constants.ErrInternal
	}

	tokenResponse, err := jwtService.GenerateToken(ctx, &jwt.GenerateTokenRequest{
		Subject: userId,
	})
	if err != nil {
		log.Error("JWT service failed to generate access token: %v", err)
		return nil, constants.ErrInternal
	}

	return &login.LoginResponse{
		AccessToken: tokenResponse.Token,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)
	jwtService = jwt.NewJwtServiceClient(network.NewGRPCClientConnection(config.GetUint("JWT_PORT"), log))
	db := mongodb.Connect(config.GetString("LOGIN_DB_HOST"), config.GetString("LOGIN_DB_USERNAME"), config.GetString("LOGIN_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("LOGIN_DB_DATABASE"), config.GetString("LOGIN_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("LOGIN_PORT"), log)
	login.RegisterLoginServiceServer(app.InternalServer, &LoginServiceServer{})
	app.Start()
}
