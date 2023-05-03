package main

import (
	"context"

	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	change_password "github.com/xd-Abi/moxie/pkg/proto/change-password"
	"github.com/xd-Abi/moxie/pkg/proto/jwt"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	log          = logging.New()
	jwtService   jwt.JwtServiceClient
	dbCollection *mongodb.MongoCollection
)

type ChangePasswordServiceServer struct {
	change_password.UnimplementedChangePasswordServiceServer
}

func (s *ChangePasswordServiceServer) ChangePassword(ctx context.Context, request *change_password.ChangePasswordRequest) (*change_password.ChangePasswordResponse, error) {
	if utils.IsEmptyOrWhitespace(request.AccessToken) {
		return nil, constants.ErrJwtEmpty
	}
	if utils.IsEmptyOrWhitespace(request.CurrentPassword) {
		return nil, constants.ErrCurrentPasswordEmpty
	}
	if utils.IsEmptyOrWhitespace(request.NewPassword) {
		return nil, constants.ErrNewPasswordEmpty
	}

	verificationResponse, err := jwtService.VerifyToken(ctx, &jwt.VerifyTokenRequest{Token: request.AccessToken})
	if err != nil {
		return nil, constants.ErrUnauthorized
	}

	user, err := dbCollection.FindOne(bson.D{{Key: "id", Value: verificationResponse.Payload["sub"]}})
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	if currentHashedPassword, ok := user["password"].(string); ok {
		if !utils.ComparePasswords([]byte(currentHashedPassword), []byte(request.CurrentPassword)) {
			return nil, constants.ErrCurrentPasswordInvalid
		}
	} else {
		log.Error("Failed to convert current password into string")
		return nil, constants.ErrInternal
	}

	newHashedPassword, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		log.Error("Failed to hash new password")
		return nil, constants.ErrInternal
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "password", Value: newHashedPassword},
		}},
	}

	_, err = dbCollection.UpdateOne(bson.D{{Key: "id", Value: verificationResponse.Payload["sub"]}}, update)
	if err != nil {
		log.Error("Failed to update user password")
		return nil, constants.ErrInternal
	}

	return &change_password.ChangePasswordResponse{
		PasswordChanged: true,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)
	jwtService = jwt.NewJwtServiceClient(network.NewGRPCClientConnection(config.GetUint("JWT_PORT"), log))
	db := mongodb.Connect(config.GetString("CHANGEPASSWORD_DB_HOST"), config.GetString("CHANGEPASSWORD_DB_USERNAME"), config.GetString("CHANGEPASSWORD_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("CHANGEPASSWORD_DB_DATABASE"), config.GetString("CHANGEPASSWORD_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("CHANGEPASSWORD_PORT"), log)
	change_password.RegisterChangePasswordServiceServer(app.InternalServer, &ChangePasswordServiceServer{})
	app.Start()
}
