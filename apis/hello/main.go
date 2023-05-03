package main

import (
	"context"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/constants"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/hello"
	"github.com/xd-Abi/moxie/pkg/utils"
)

var (
	log = logging.New()
)

type HelloServiceServer struct {
	SendGridApiKey string
	SendGridSender string
	hello.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) SendWelcomeEmail(ctx context.Context, request *hello.WelcomeEmailRequest) (*hello.WelcomeEmailResponse, error) {
	if utils.IsEmptyOrWhitespace(request.Email) {
		return nil, constants.ErrEmailEmpty
	}
	if !utils.IsEmail(request.Email) {
		return nil, constants.ErrEmailInvalid
	}
	if utils.IsEmptyOrWhitespace(request.Username) {
		return nil, constants.ErrUsernameEmpty
	}

	from := mail.NewEmail("Moxie", s.SendGridSender)
	subject := "Welcome to moxie 😎"
	to := mail.NewEmail(request.Username, request.Email)

	// @TODO: Load html content from e.g. template.html file
	plainTextContent := fmt.Sprintf("Hi %s", request.Username)
	htmlContent := fmt.Sprintf("Hi <strong>%s</strong>", request.Username)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.SendGridApiKey)
	_, err := client.Send(message)

	if err != nil {
		log.Error("Failed to send email: %v", err)
		return nil, constants.ErrInternal
	} else {
		log.Info("Send welcome email to %v", request.Email)
	}

	return &hello.WelcomeEmailResponse{
		Success: true,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)

	app := network.NewMicroServiceServer(config.GetUint("HELLO_PORT"), log)
	hello.RegisterHelloServiceServer(app.InternalServer, &HelloServiceServer{
		SendGridApiKey: config.GetString("HELLO_SENDGRID_API_KEY"),
		SendGridSender: config.GetString("HELLO_SENDGRID_SENDER"),
	})
	app.Start()
}
