package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/hello"
	"github.com/xd-Abi/moxie/pkg/proto/jwt"
	"github.com/xd-Abi/moxie/pkg/proto/signup"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	log          = logging.New()
	jwtService   jwt.JwtServiceClient
	helloService hello.HelloServiceClient
	dbCollection *mongodb.MongoCollection
)

type SignUpServiceServer struct {
	signup.UnimplementedSignUpServiceServer
}

func (s *SignUpServiceServer) SignUp(ctx context.Context, request *signup.SignUpRequest) (*signup.SignUpResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Username) {
		return nil, constants.ErrUsernameEmpty
	}
	if utils.IsEmptyOrWhitespace(request.Email) {
		return nil, constants.ErrEmailEmpty
	}
	if !utils.IsEmail(request.Email) {
		return nil, constants.ErrEmailInvalid
	}
	if utils.IsEmptyOrWhitespace(request.Password) {
		return nil, constants.ErrPasswordEmpty
	}
	if _, err := dbCollection.FindOne(bson.D{{Key: "email", Value: request.Email}}); err == nil {
		return nil, constants.ErrEmailAlreadyExists
	}
	if _, err := dbCollection.FindOne(bson.D{{Key: "username", Value: request.Username}}); err == nil {
		return nil, constants.ErrUsernameAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Error("Failed to hash password: %v", err)
		return nil, constants.ErrInternal
	}

	user := bson.D{
		{Key: "id", Value: utils.GenerateUUID()},
		{Key: "username", Value: request.Username},
		{Key: "email", Value: request.Email},
		{Key: "password", Value: hashedPassword},
	}

	_, err = dbCollection.InsertOne(user)
	if err != nil {
		return nil, constants.ErrInternal
	}

	return nil, nil
}

func main() {
	config.LoadEnvVariables(log)
	jwtService = jwt.NewJwtServiceClient(network.NewGRPCClientConnection(config.GetUint("JWT_PORT"), log))
	helloService = hello.NewHelloServiceClient(network.NewGRPCClientConnection(config.GetUint("HELLO_PORT"), log))

	db := mongodb.Connect("mongodb://localhost:27001", "auth-api", "VpGdxrtBbzs0Zgxn5o", log)
	dbCollection = db.GetCollection("auth", "users")

	app := network.NewMicroServiceServer(config.GetUint("SIGN_UP_PORT"), log)
	signup.RegisterSignUpServiceServer(app.InternalServer, &SignUpServiceServer{})
	app.Start()
}
