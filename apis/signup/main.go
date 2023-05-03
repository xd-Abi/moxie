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

	userId := utils.GenerateUUID()
	user := bson.D{
		{Key: "id", Value: userId},
		{Key: "username", Value: request.Username},
		{Key: "email", Value: request.Email},
		{Key: "password", Value: hashedPassword},
	}

	_, err = dbCollection.InsertOne(user)
	if err != nil {
		return nil, constants.ErrInternal
	}

	_, err = helloService.SendWelcomeEmail(ctx, &hello.WelcomeEmailRequest{
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		log.Error("Hello service failed to send welcome email: %v", err)
	}

	tokenResponse, err := jwtService.GenerateToken(ctx, &jwt.GenerateTokenRequest{
		Subject: userId,
	})
	if err != nil {
		log.Error("JWT service failed to generate access token: %v", err)
		return nil, constants.ErrInternal
	}

	return &signup.SignUpResponse{
		AccessToken: tokenResponse.Token,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)
	jwtService = jwt.NewJwtServiceClient(network.NewGRPCClientConnection(config.GetUint("JWT_PORT"), log))
	helloService = hello.NewHelloServiceClient(network.NewGRPCClientConnection(config.GetUint("HELLO_PORT"), log))

	db := mongodb.Connect(config.GetString("SIGNUP_DB_HOST"), config.GetString("SIGNUP_DB_USERNAME"), config.GetString("SIGNUP_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("SIGNUP_DB_DATABASE"), config.GetString("SIGNUP_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("SIGNUP_PORT"), log)
	signup.RegisterSignUpServiceServer(app.InternalServer, &SignUpServiceServer{})
	app.Start()
}
