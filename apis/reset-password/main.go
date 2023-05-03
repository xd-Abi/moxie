package main

import (
	"context"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/mongodb"
	"github.com/xd-Abi/moxie/pkg/network"
	reset_password "github.com/xd-Abi/moxie/pkg/proto/reset-password"
	"github.com/xd-Abi/moxie/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	log          = logging.New()
	dbCollection *mongodb.MongoCollection
)

type ResetPasswordServiceServer struct {
	SendGridApiKey string
	SendGridSender string
	reset_password.UnimplementedResetPasswordServiceServer
}

// @TODO: it is possible to reset password infinitely
func (s *ResetPasswordServiceServer) ResetPassword(ctx context.Context, request *reset_password.ResetPasswordRequest) (*reset_password.ResetPasswordResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Email) {
		return nil, constants.ErrEmailEmpty
	}
	if !utils.IsEmail(request.Email) {
		return nil, constants.ErrEmailInvalid
	}

	user, err := dbCollection.FindOne(bson.D{{Key: "email", Value: request.Email}})
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	username, ok := user["username"].(string)
	if !ok {
		log.Error("Failed to convert username into string")
		return nil, constants.ErrInternal
	}

	newPassword := utils.GenerateUUID()

	from := mail.NewEmail("Moxie", s.SendGridSender)
	subject := "Password Reset - Moxie ⚡"
	to := mail.NewEmail(username, request.Email)

	// @TODO: Load html content from e.g. template.html file
	plainTextContent := fmt.Sprintf("Hi %s, here is your new password: %v", username, newPassword)
	htmlContent := fmt.Sprintf("Hi <strong>%s</strong>, here is your new password: %v", username, newPassword)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.SendGridApiKey)
	_, err = client.Send(message)

	if err != nil {
		log.Error("Failed to send reset password email: %v", err)
		return nil, constants.ErrInternal
	} else {
		log.Info("Reset password email to %v", request.Email)
	}

	newHashedPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		log.Error("Failed to hash password")
		return nil, constants.ErrInternal
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "password", Value: newHashedPassword},
		}},
	}

	_, err = dbCollection.UpdateOne(bson.D{{Key: "email", Value: request.Email}}, update)
	if err != nil {
		log.Error("Failed to update user's password")
		return nil, constants.ErrInternal
	}

	return &reset_password.ResetPasswordResponse{
		PasswordReset: true,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)
	db := mongodb.Connect(config.GetString("RESETPASSWORD_DB_HOST"), config.GetString("RESETPASSWORD_DB_USERNAME"), config.GetString("RESETPASSWORD_DB_PASSWORD"), log)
	dbCollection = db.GetCollection(config.GetString("RESETPASSWORD_DB_DATABASE"), config.GetString("RESETPASSWORD_DB_COLLECTION"))

	app := network.NewMicroServiceServer(config.GetUint("RESETPASSWORD_PORT"), log)
	reset_password.RegisterResetPasswordServiceServer(app.InternalServer, &ResetPasswordServiceServer{
		SendGridApiKey: config.GetString("HELLO_SENDGRID_API_KEY"),
		SendGridSender: config.GetString("HELLO_SENDGRID_SENDER"),
	})
	app.Start()
}
