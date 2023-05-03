package main

import (
	"context"
	"fmt"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/jwt"
	"github.com/xd-Abi/moxie/pkg/utils"
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
	if utils.IsEmptyOrWhitespace(request.Subject) {
		return nil, constants.ErrSubjectEmpty
	}

	claims := jwtGo.MapClaims{
		"sub": request.Subject,
		"exp": time.Now().Add(s.TokenExpiration).Unix(),
	}

	plainToken := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	signedToken, err := plainToken.SignedString([]byte(s.Secret))

	if err != nil {
		log.Error("Failed to sign JSON web token: %v", err)
		return nil, constants.ErrInternal
	}

	return &jwt.GenerateTokenResponse{Token: signedToken}, nil
}

func (s *JwtServiceServer) VerifyToken(ctx context.Context, request *jwt.VerifyTokenRequest) (*jwt.VerifyTokenResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Token) {
		return nil, constants.ErrJwtEmpty
	}

	parsedToken, err := jwtGo.ParseWithClaims(request.Token, jwtGo.MapClaims{}, func(token *jwtGo.Token) (interface{}, error) {
		// Verify the signature of the token
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, constants.ErrInternal
		}

		return []byte(s.Secret), nil
	})

	if err != nil {
		return nil, constants.ErrJwtInvalid
	}

	claims, ok := parsedToken.Claims.(jwtGo.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, constants.ErrJwtInvalid
	}

	payload := make(map[string]string)
	for key, value := range claims {
		payload[key] = fmt.Sprintf("%v", value)
	}

	return &jwt.VerifyTokenResponse{
		Payload: payload,
	}, nil
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
