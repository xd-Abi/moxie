package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/auth"
	"github.com/xd-Abi/moxie/pkg/rabbitmq"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Id       string `bson:"id,omitempty"`
	Username string `bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type RefreshToken struct {
	UserId     string `bson:"user_id,omitempty"`
	Token      string `bson:"token,omit"`
	Expiration int64  `bson:"expiration,omit"`
}

type MongoConfig struct {
	Uri                    string
	Name                   string
	Username               string
	Password               string
	UserCollection         string
	RefreshTokenCollection string
}

type Config struct {
	JwtSigningKey          []byte
	JwtExpiration          time.Duration
	RefreshTokenExpiration time.Duration
	Mongo                  *MongoConfig
	RabbitMQUrl            string
}

type AuthServiceServer struct {
	Log                    *logging.Log
	Config                 *Config
	userCollection         *mongodb.MongoCollection
	refreshTokenCollection *mongodb.MongoCollection
	rabbitMQConnection     *rabbitmq.RabbitMQConnection

	auth.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(config *Config, log *logging.Log) *AuthServiceServer {
	db := mongodb.Connect(config.Mongo.Uri, config.Mongo.Username, config.Mongo.Password, log)
	userCollection := db.GetCollection(config.Mongo.Name, config.Mongo.UserCollection)
	refreshTokenCollection := db.GetCollection(config.Mongo.Name, config.Mongo.RefreshTokenCollection)
	rabbitMQConnection := rabbitmq.NewConnection(config.RabbitMQUrl, log)
	rabbitMQConnection.DeclareExchange(rabbitmq.AuthExchangeKey)
	rabbitMQConnection.DeclareQueue(rabbitmq.ProfileQueueKey)
	rabbitMQConnection.Bind(rabbitmq.ProfileQueueKey, rabbitmq.UserSignUpEventKey, rabbitmq.AuthExchangeKey)

	return &AuthServiceServer{
		Log:                    log,
		Config:                 config,
		userCollection:         userCollection,
		refreshTokenCollection: refreshTokenCollection,
		rabbitMQConnection:     rabbitMQConnection,
	}
}

func (s *AuthServiceServer) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.SignUpResponse, error) {
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
	if _, err := s.userCollection.FindOne(bson.D{{Key: "email", Value: request.Email}}); err == nil {
		return nil, constants.ErrEmailAlreadyExists
	}
	if _, err := s.userCollection.FindOne(bson.D{{Key: "username", Value: request.Username}}); err == nil {
		return nil, constants.ErrUsernameAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		s.Log.Error("Failed to hash password: %v", err)
		return nil, constants.ErrInternal
	}

	user := User{
		Id:       utils.GenerateUUID(),
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	_, err = s.userCollection.InsertOne(user)
	if err != nil {
		s.Log.Error("Failed to insert user: %v", err)
		return nil, constants.ErrInternal
	}

	accessToken, err := s.generateAccessToken(user.Id)
	if err != nil {
		s.Log.Error("Failed to generate access token: %v", err)
		return nil, constants.ErrInternal
	}

	refreshToken, err := s.generateRefreshToken(user.Id)
	if err != nil {
		s.Log.Error("Failed to generate refresh token: %v", err)
		return nil, constants.ErrInternal
	}

	s.rabbitMQConnection.Publish(rabbitmq.AuthExchangeKey, rabbitmq.NewSignUpEvent(
		rabbitmq.UserSignUpEventPayload{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	))

	return &auth.SignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Email) {
		return nil, constants.ErrEmailEmpty
	}
	if !utils.IsEmail(request.Email) {
		return nil, constants.ErrEmailInvalid
	}
	if utils.IsEmptyOrWhitespace(request.Password) {
		return nil, constants.ErrPasswordEmpty
	}

	var user User
	err := s.userCollection.FindOneAndDecode(bson.D{{Key: "email", Value: request.Email}}, &user)
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	if !utils.ComparePasswords([]byte(user.Password), []byte(request.Password)) {
		return nil, constants.ErrPasswordInvalid
	}

	accessToken, err := s.generateAccessToken(user.Id)
	if err != nil {
		s.Log.Error("Failed to generate access token: %v", err)
		return nil, constants.ErrInternal
	}

	refreshToken, err := s.generateRefreshToken(user.Id)
	if err != nil {
		s.Log.Error("Failed to generate refresh token: %v", err)
		return nil, constants.ErrInternal
	}

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) VerifyToken(ctx context.Context, request *auth.TokenVerificationRequest) (*auth.TokenVerificationResponse, error) {
	if utils.IsEmptyOrWhitespace(request.AccessToken) {
		return nil, constants.ErrJwtEmpty
	}

	parsedToken, err := jwt.ParseWithClaims(request.AccessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the signature of the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constants.ErrInternal
		}

		return []byte(s.Config.JwtSigningKey), nil
	})

	if err != nil {
		return nil, constants.ErrJwtInvalid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, constants.ErrJwtInvalid
	}

	payload := make(map[string]string)
	for key, value := range claims {
		payload[key] = fmt.Sprintf("%v", value)
	}

	return &auth.TokenVerificationResponse{
		Payload: payload,
	}, nil
}

func (s *AuthServiceServer) RefreshToken(ctx context.Context, request *auth.TokenRefreshRequest) (*auth.TokenRefreshResponse, error) {
	if utils.IsEmptyOrWhitespace(request.RefreshToken) {
		return nil, constants.ErrRefreshTokenEmpty
	}

	var refreshToken RefreshToken
	err := s.refreshTokenCollection.FindOneAndDecode(bson.D{{Key: "token", Value: request.RefreshToken}}, &refreshToken)

	if err != nil {
		return nil, constants.ErrRefreshTokenInvalid
	}

	if time.Now().Unix() > refreshToken.Expiration {
		return nil, constants.ErrRefreshTokenInvalid
	}

	accessToken, err := s.generateAccessToken(refreshToken.UserId)
	if err != nil {
		s.Log.Error("Failed to generate access token: %v", err)
		return nil, constants.ErrInternal
	}

	return &auth.TokenRefreshResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthServiceServer) GetHealth(ctx context.Context, request *auth.HealthRequest) (*auth.HealthResponse, error) {
	return &auth.HealthResponse{
		Message:   "I am up and running",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (s *AuthServiceServer) generateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(s.Config.JwtExpiration).Unix(),
	}

	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := plainToken.SignedString([]byte(s.Config.JwtSigningKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthServiceServer) generateRefreshToken(userId string) (string, error) {
	var refreshToken RefreshToken
	refreshToken.Expiration = time.Now().Add(s.Config.RefreshTokenExpiration).Unix()
	refreshToken.UserId = userId
	refreshToken.Token = utils.GenerateUUID()

	_, err := s.refreshTokenCollection.FindOne(bson.D{{Key: "user_id", Value: userId}})

	if err != nil {
		_, err := s.refreshTokenCollection.InsertOne(refreshToken)

		if err != nil {
			s.Log.Error("Failed to insert refresh token: %v", err)
			return "", err
		}
	} else {
		_, err := s.refreshTokenCollection.UpdateOne(bson.D{{Key: "user_id", Value: userId}}, bson.M{"$set": refreshToken})

		if err != nil {
			s.Log.Error("Failed to update refresh token: %v", err)
			return "", err
		}
	}

	return refreshToken.Token, nil
}

func main() {
	log := logging.New()
	config.LoadEnvVariables(log)
	app := network.NewMicroServiceServer(config.GetUint("AUTH_PORT"), log)
	config := Config{
		JwtSigningKey:          []byte(config.GetString("AUTH_JWT_SIGNING_KEY")),
		JwtExpiration:          time.Duration(config.GetUint("AUTH_JWT_EXPIRATION")),
		RefreshTokenExpiration: time.Duration(config.GetUint("AUTH_REFRESH_TOKEN_EXPIRATION")),
		RabbitMQUrl:            config.GetString("AUTH_RABBITMQ_URL"),
		Mongo: &MongoConfig{
			Uri:                    config.GetString("AUTH_DB_URI"),
			Username:               config.GetString("AUTH_DB_USERNAME"),
			Password:               config.GetString("AUTH_DB_PASSWORD"),
			Name:                   config.GetString("AUTH_DB_DATABASE"),
			UserCollection:         config.GetString("AUTH_DB_USER_COLLECTION"),
			RefreshTokenCollection: config.GetString("AUTH_DB_REFRESH_TOKEN_COLLECTION"),
		},
	}

	auth.RegisterAuthServiceServer(app.InternalServer, NewAuthServiceServer(&config, log))
	app.Start()
}
